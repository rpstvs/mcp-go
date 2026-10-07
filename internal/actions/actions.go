package actions

import (
	"context"

	"github.com/rpstvs/mcp-go/types"
)

type Handler func(ctx context.Context, run *types.Run)
type Action struct {
	handler Handler
}
