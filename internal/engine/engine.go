package engine

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
	"github.com/rpstvs/mcp-go/internal/workflow"
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

func (e *Engine) RegisterWorkflows(wfs []workflow.Workflow) {
	for _, wf := range wfs {
		e.registry.Register(wf.Name, wf)
	}
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

	wf, ok := e.registry.Get(kind)
	if !ok {
		return nil, fmt.Errorf("unknown workflow ")
	}

	e.Store.CreateRun(run)

	ctx := context.Background()
	e.RunWorkflow(ctx, run, wf)

	return run, nil
}

func (e *Engine) RunWorkflow(ctx context.Context, run *types.Run, workflow workflow.Workflow) (*types.Run, error) {

	run.State = types.StateRunning

	session, err := e.Sessions.Create(ctx)
	if err != nil {
		run.State = types.StateFailed
		return run, err
	}

	run.Session = session
	var output string
	for _, StepHandler := range workflow.Steps {
		output, err = StepHandler(ctx, run)

		if err != nil {
			run.State = types.StateFailed
			return run, err
		}
	}

	run.OutputJSON = output

	run.State = types.StateCompleted
	return run, nil
}

func (e *Engine) ResumeWorkflow() {

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
