package workflow

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"
	"uuid"

	"github.com/rpstvs/mcp-go/config"
)

type Engine struct {
	config   config.Config
	registry *Registry
}

func NewEngine(cfg config.Config, reg *Registry) *Engine {
	e := &Engine{
		config:   cfg,
		registry: NewRegistry(),
	}

	e.registry.Register(FEATURE_IMPLEMENTATION_WORKFLOW)
	e.registry.Register(PR_REVIEW_WORKFLOW)

	return e
}

func (e *Engine) Start(kind WorkflowType, workspace string, input any) (*Run, error) {
	workspace, err := filepath.Abs(workspace)

	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	run := &Run{
		ID:        uuid.NewString(),
		Type:      kind,
		State:     StateQueued,
		Workspace: workspace,
		InputJSON: string(raw),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if ok := e.registry.Get(kind); !ok {
		return nil, fmt.Errorf("unknown workflow ")
	}

	go e.execute(run.ID)

	return run, nil
}

func (e *Engine) execute(runid string) error {

}
