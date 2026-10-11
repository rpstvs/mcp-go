package engine

import (
	"github.com/rpstvs/mcp-go/internal/workflow"
	"github.com/rpstvs/mcp-go/types"
)

type Registry struct {
	workflows map[types.WorkflowType]workflow.Workflow
}

func NewRegistry() *Registry {
	return &Registry{
		workflows: make(map[types.WorkflowType]workflow.Workflow),
	}
}

func (r *Registry) Register(kind types.WorkflowType, wf workflow.Workflow) {
	r.workflows[kind] = wf
}

func (r *Registry) Get(kind types.WorkflowType) (workflow.Workflow, bool) {
	val, ok := r.workflows[kind]

	return val, ok
}
