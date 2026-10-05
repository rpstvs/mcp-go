package workflow

import "time"

type WorkflowType string

const (
	WorkflowReviewPR        WorkflowType = "review_pr"
	WorkflowImplementTicket WorkflowType = "implement_ticket"
)

type RunState string

const (
	StateQueued        RunState = "queued"
	StateRunning       RunState = "running"
	StateAwaitingPlan  RunState = "awaiting_plan"
	StateReviewing     RunState = "reviewing"
	StateVerifying     RunState = "verifying"
	StateAwaitingFinal RunState = "awaiting_final"
	StateCompleted     RunState = "completed"
	StateRejected      RunState = "rejected"
	StateFailed        RunState = "failed"
)

type Run struct {
	ID         string       `json:"id"`
	Type       WorkflowType `json:"type"`
	State      RunState     `json:"state"`
	Workspace  string       `json:"workspace"`
	InputJSON  string       `json:"input_json"`
	OutputJSON string       `json:"output_json,omitempty"`
	Error      string       `json:"error,omitempty"`
	Approval   string       `json:"approval,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type ReviewFinding struct {
	Reviewer string `json:"reviewer"`
	Severity string `json:"severity"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
	Message  string `json:"message"`
}

type PRReviewInput struct {
	Workspace   string `json:"workspace" jsonschema:"absolute path to the repository"`
	PullRequest int    `json:"pull_request" jsonschema:"GitHub pull request number"`
}

type TicketInput struct {
	Workspace string `json:"workspace" jsonschema:"absolute path to the repository"`
	TicketID  string `json:"ticket_id" jsonschema:"Azure DevOps work item ID"`
}
