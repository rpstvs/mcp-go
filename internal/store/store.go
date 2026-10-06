package store

import "github.com/rpstvs/mcp-go/types"

type Store interface {
	CreateRun(run *types.Run)
	GetRun(runid string) *types.Run
	UpdateRun(run *types.Run)
	DeleteRun(runid string)
}
