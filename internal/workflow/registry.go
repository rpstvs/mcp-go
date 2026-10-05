package workflow

type Handler func(run *Run) error

type Registry struct {
	handlers map[WorkflowType]Handler
}

func NewRegistry() *Registry {
	return &Registry{
		handlers: make(map[WorkflowType]Handler),
	}
}

func (r *Registry) Register(kind WorkflowType, handler Handler) {
	r.handlers[kind] = handler
}

func (r *Registry) Get(kind WorkflowType) (Handler, bool) {
	val, ok := r.handlers[kind]

	return val, ok
}
