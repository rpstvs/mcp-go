package actions

import (
	"context"

	"github.com/rpstvs/mcp-go/types"
)

type Handler func(ctx context.Context, run *types.Run) (string, error)

type Actions struct {
	ReviewPrAction Handler
	HumanGate      Handler
}

func NewActions() *Actions {
	return &Actions{
		ReviewPrAction: ReviewPrAction(),
		HumanGate:      HumanGate(),
	}
}
