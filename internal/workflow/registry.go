package workflow

import "github.com/rpstvs/mcp-go/types"

type HandlerFactory func() (Workflow, error)

type Registry struct {
	workflows map[types.WorkflowType]HandlerFactory
}

func NewRegistry() *Registry {
	return &Registry{
		workflows: make(map[types.WorkflowType]HandlerFactory),
	}
}

func (r *Registry) Register(kind types.WorkflowType, hf HandlerFactory) {
	r.workflows[kind] = hf
}

func (r *Registry) Get(kind types.WorkflowType) (HandlerFactory, bool) {
	val, ok := r.workflows[kind]

	return val, ok
}
