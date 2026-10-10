package workflow

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/mcp-go/config"
	"github.com/rpstvs/mcp-go/internal/copilot"
	"github.com/rpstvs/mcp-go/internal/store"
	"github.com/rpstvs/mcp-go/types"
)

type Engine struct {
	config   config.Config
	registry *Registry
	Store    store.Store
	Sessions copilot.SessionFactory
}

func NewEngine(cfg config.Config, store store.Store, factory copilot.SessionFactory) *Engine {
	e := &Engine{
		config:   cfg,
		registry: NewRegistry(),
		Store:    store,
		Sessions: factory,
	}

	return e
}

func (e *Engine) Start(kind types.WorkflowType, workspace string, input any) (*types.Run, error) {
	workspace, err := filepath.Abs(workspace)

	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	run := &types.Run{
		ID:        uuid.NewString(),
		Type:      kind,
		State:     types.StateQueued,
		Workspace: workspace,
		InputJSON: string(raw),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if _, ok := e.registry.Get(kind); !ok {
		return nil, fmt.Errorf("unknown workflow ")
	}

	e.Store.CreateRun(run)

	go e.execute(run.ID)

	ctx := context.Background()
	go e.RunWorkflow(ctx, run, workflow)

	return run, nil
}

func (e *Engine) RunWorkflow(ctx context.Context, run *types.Run, workflow Workflow) error {

	session, err := e.Sessions.Create(ctx)
	if err != nil {
		return err
	}

	run.Session = session

	for _, StepHandler := range workflow.Steps {
		StepHandler(ctx, run)
	}
	return nil
}

func (e *Engine) execute(runid string) error {

	run := e.Store.GetRun(runid)

	run.State = types.StateRunning

	handler, ok := e.registry.Get(run.Type)

	if !ok {
		return fmt.Errorf("handler not registered")
	}

	handler(run)

	e.Store.UpdateRun(run)

	return nil
}

func (e *Engine) Approve(id string, comment string) error {

	run := e.Store.GetRun(id)

	run.Approval = "approved: " + comment
	run.State = types.StateRunning

	e.Store.UpdateRun(run)

	return nil
}

func (e *Engine) Reject(id, reason string) error {
	run := e.Store.GetRun(id)

	run.Approval = "rejected: " + reason
	run.State = types.StateRejected

	e.Store.UpdateRun(run)

	return nil
}

func (e *Engine) runPrReview(run *types.Run) error {
	var input types.PRReviewInput

	if err := json.Unmarshal([]byte(run.InputJSON), &input); err != nil {
		return err
	}

	ctx := context.Background()

	reviewer := copilot.NewCopilotReviewer(e.config)

	result, err := reviewer.ReviewPR(ctx, input.Workspace, input.PullRequest)

	if err != nil {
		return err
	}

	run.OutputJSON = result
	run.State = types.StateCompleted

	return nil
}
