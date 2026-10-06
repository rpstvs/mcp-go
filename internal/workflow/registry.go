package workflow

import "github.com/rpstvs/mcp-go/types"

type Handler func(run *types.Run) error

type Registry struct {
	handlers map[types.WorkflowType]Handler
}

func NewRegistry() *Registry {
	return &Registry{
		handlers: make(map[types.WorkflowType]Handler),
	}
}

func (r *Registry) Register(kind types.WorkflowType, handler Handler) {
	r.handlers[kind] = handler
}

func (r *Registry) Get(kind types.WorkflowType) (Handler, bool) {
	val, ok := r.handlers[kind]

	return val, ok
}
