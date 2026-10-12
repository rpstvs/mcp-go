package actions

import (
	"context"

	"github.com/rpstvs/mcp-go/types"
)

func HumanGate(ctx context.Context, run *types.Run) (*types.Run, error) {
	run.State = types.StateAwaitingDecision

}
