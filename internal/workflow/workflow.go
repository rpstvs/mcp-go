package workflow

import (
	"github.com/rpstvs/mcp-go/internal/actions"
)

type Workflow struct {
	Steps []actions.Handler
}

func ReviewPrWorkflow(a actions.Actions) Workflow {
	return Workflow{
		Steps: []actions.Handler{
			a.GetPrAction,
			a.ReviewPrAction,
		},
	}
}
