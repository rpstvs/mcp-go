package workflow

import "github.com/rpstvs/mcp-go/types"

type Registry struct {
	workflows map[types.WorkflowType]Workflow
}

func NewRegistry() *Registry {
	return &Registry{
		workflows: make(map[types.WorkflowType]Workflow),
	}
}

func (r *Registry) Register(kind types.WorkflowType, wf Workflow) {
	r.workflows[kind] = wf
}

func (r *Registry) Get(kind types.WorkflowType) (Workflow, bool) {
	val, ok := r.workflows[kind]

	return val, ok
}
