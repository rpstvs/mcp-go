package actions

import (
	"context"

	"github.com/rpstvs/mcp-go/types"
)

type Handler func(ctx context.Context, run *types.Run)

type Actions struct {
	GetPrAction    Handler
	ReviewPrAction Handler
}

func NewActions(copilot *sdk) *Actions {
	return &Actions{
		GetPrAction:    GetPrAction(copilot),
		ReviewPrAction: ReviewPrAction(copilot),
	}
}
