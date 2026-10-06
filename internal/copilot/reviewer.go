package copilot

import (
	"context"
	"encoding/json"
	"fmt"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/rpstvs/mcp-go/config"
)

type CopilotReviewer struct {
	cfg config.Config
}

func NewCopilotReviewer(cfg config.Config) *CopilotReviewer {
	return &CopilotReviewer{
		cfg: cfg,
	}
}

type reviewResult struct {
	PR      int    `json:"pull_request"`
	Summary string `json:"summary"`
	Raw     string `json:"raw"`
}

func (r *CopilotReviewer) ReviewPR(ctx context.Context, workspace string, pr int) (string, error) {
	client := copilot.NewClient(&copilot.ClientOptions{
		LogLevel: "error",
	})

	if err := client.Start(ctx); err != nil {
		return "", nil
	}

	defer client.Stop()

	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		WorkingDirectory: workspace,
	})

	if err != nil {
		return "", err
	}

	defer session.Disconnect()

	doneCh := make(chan string, 1)
	errCh := make(chan error, 1)

	session.On(func(event copilot.SessionEvent) {
		if data, ok := event.Data.(*copilot.AssistantMessageData); ok {
			select {
			case doneCh <- data.Content:
			default:
			}
		}
	})

	_, err = session.SendAndWait(ctx, copilot.MessageOptions{
		Prompt: fmt.Sprintf("Review the pr %d", pr),
	})

	if err != nil {
		return "", err
	}

	select {
	case response := <-doneCh:
		out, _ := json.Marshal(reviewResult{
			PR:      pr,
			Summary: "PR review completed",
			Raw:     response,
		})
		return string(out), nil
	case err := <-errCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
