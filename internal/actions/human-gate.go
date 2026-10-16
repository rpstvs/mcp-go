package actions

import (
	"context"
	"errors"

	"github.com/rpstvs/mcp-go/types"
)

var (
	ErrHumanRejected = errors.New("Rejected at human gate")
	ErrHumanApproved = errors.New("Acepted at human gate")
)

func HumanGate() Handler {
	return func(ctx context.Context, run *types.Run) (string, error) {

		switch {
		case run.Decision == types.DecisionRejected:
			run.State = types.StateAwaitingDecision
			return "", ErrHumanRejected
		case run.Decision == types.DecisionApproved:
			return "Approved", nil
		default:
			run.Decision = types.AwaitingDecision
			return "Awaiting decision", nil
		}

	}
}
