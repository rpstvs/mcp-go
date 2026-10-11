package actions

import (
	"context"
	"encoding/json"
	"fmt"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/rpstvs/mcp-go/types"
)

type reviewResult struct {
	pr       int
	result   string
	response any
}

func ReviewPrAction() Handler {
	return func(ctx context.Context, run *types.Run) (string, error) {

		doneCh := make(chan string, 1)
		errCh := make(chan error, 1)

		run.Session.On(func(event copilot.SessionEvent) {
			if data, ok := event.Data.(*copilot.AssistantMessageData); ok {
				select {
				case doneCh <- data.Content:
				default:
				}

			}
		})

		_, err := run.Session.SendAndWait(ctx, copilot.MessageOptions{
			Prompt: fmt.Sprintf(`Review the PR %s`, run.InputJSON),
		})

		if err != nil {
			return "", err
		}

		select {
		case response := <-doneCh:

			out, _ := json.Marshal(reviewResult{
				pr:       2,
				result:   "pr completed",
				response: response,
			})

			return string(out), nil
		case err := <-errCh:
			return "", err
		case <-ctx.Done():
			return "", ctx.Err()
		}

	}
}
