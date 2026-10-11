package workflow

import (
	"github.com/rpstvs/mcp-go/internal/actions"
	"github.com/rpstvs/mcp-go/types"
)

type Workflow struct {
	Name  types.WorkflowType
	Steps []actions.Handler
}

func CreateWorkflows(a *actions.Actions) []Workflow {
	return []Workflow{
		ReviewPrWorkflow(a),
	}
}

func ReviewPrWorkflow(a *actions.Actions) Workflow {
	return Workflow{
		Name: types.WorkflowReviewPR,
		Steps: []actions.Handler{
			a.ReviewPrAction,
		},
	}
}
