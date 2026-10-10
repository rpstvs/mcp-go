package actions

import (
	"context"
	"fmt"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/rpstvs/mcp-go/types"
)

func ReviewPrAction() Handler {
	return func(ctx context.Context, run *types.Run) {
		run.Session.Send(ctx, copilot.MessageOptions{
			Prompt: fmt.Sprintf(`Review the PR %d`, run.InputJSON),
		})

	}

}
