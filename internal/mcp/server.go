package mcp

import (
	"context"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rpstvs/mcp-go/internal/workflow"
	"github.com/rpstvs/mcp-go/types"
)

func Register(server *sdk.Server, engine *workflow.Engine) {
	registerStartWorkflow(server, engine)
}

type StartWorkflowInput struct {
	Type      string         `json:"type" jsonschema:"workflow type, for example review_pr or implement_ticket"`
	Workspace string         `json:"workspace" jsonschema:"absolute repository path"`
	Input     map[string]any `json:"input" jsonschema:"workflow-specific input"`
}

type StartWorkflowOutput struct {
	RunID string `json:"run_id"`
	Type  string `json:"type"`
	State string `json:"state"`
}

func registerStartWorkflow(server *sdk.Server, engine *workflow.Engine) {
	sdk.AddTool(
		server,
		&sdk.Tool{
			Name:        "start_workflow",
			Description: "Start an engineering workflow on a repo",
		},

		func(ctx context.Context, req *sdk.CallToolRequest, in StartWorkflowInput) (*sdk.CallToolResult, StartWorkflowOutput, error) {

			run, err := engine.Start(types.WorkflowType(in.Type), in.Workspace, in.Input)

			if err != nil {
				return nil, StartWorkflowOutput{}, err
			}
			return nil, StartWorkflowOutput{
				RunID: run.ID,
				Type:  string(run.Type),
				State: string(run.State),
			}, nil
		},
	)
}
