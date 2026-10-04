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

func (r *Registry) GetHandler(kind WorkflowType) Handler {
	val, ok := r.handlers[kind]

	if !ok {
		return nil
	}
	return val
}
