package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

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
	i := +1
	run := &types.Run{
		ID:        fmt.Sprintf("%s-%d", kind, i),
		Type:      kind,
		State:     types.StateQueued,
		Workspace: workspace,
		InputJSON: string(raw),
		CreatedAt: now,
		UpdatedAt: now,
		StepIndex: 0,
		Decision:  types.AwaitingDecision,
	}

	wf, ok := e.registry.Get(kind)
	if !ok {
		return nil, fmt.Errorf("unknown workflow ")
	}
	run.TotalSteps = len(wf.Steps)
	e.Store.CreateRun(run)

	ctx := context.Background()

	session, err := e.Sessions.Create(ctx)
	if err != nil {
		run.State = types.StateFailed
		return run, err
	}

	run.Session = session

	run, err = e.RunWorkflow(ctx, run, wf)

	if err != nil {
		return nil, err
	}

	return run, nil
}

func (e *Engine) RunWorkflow(ctx context.Context, run *types.Run, workflow workflow.Workflow) (*types.Run, error) {

	run.State = types.StateRunning

	var output string
	var err error
	for run.StepIndex < len(workflow.Steps) {
		output, err = workflow.Steps[run.StepIndex](ctx, run)

		if err != nil {
			run.State = types.StateFailed
			run.Error = err.Error()
			return run, err
		}
		run.StepIndex++
		if run.State == types.StateAwaitingDecision {
			return run, nil
		}
	}

	run.OutputJSON = output

	run.State = types.StateCompleted
	return run, nil
}

func (e *Engine) Approve(id string, comment string) (*types.Run, error) {

	run := e.Store.GetRun(id)

	if run.Decision != types.AwaitingDecision {
		return nil, fmt.Errorf("run not awaiting decision")
	}

	run.Approval = "approved: " + comment
	run.Decision = types.DecisionApproved
	wf, exists := e.registry.workflows[run.Type]

	if !exists {
		return nil, fmt.Errorf("workflow doesnt exist.")
	}

	e.Store.UpdateRun(run)
	run, err := e.RunWorkflow(context.Background(), run, wf)

	if err != nil {
		return nil, err
	}

	return run, nil
}

func (e *Engine) Reject(id, reason string) (*types.Run, error) {
	run := e.Store.GetRun(id)

	if run.Decision != types.AwaitingDecision {
		return nil, fmt.Errorf("run not awaiting decision")
	}

	run.Approval = "rejected: " + reason
	run.Decision = types.DecisionRejected

	wf, exists := e.registry.workflows[run.Type]

	if !exists {
		return nil, fmt.Errorf("workflow not registered")
	}

	e.Store.UpdateRun(run)

	run, err := e.RunWorkflow(context.Background(), run, wf)

	if err != nil {
		return nil, err
	}

	return run, nil
}
