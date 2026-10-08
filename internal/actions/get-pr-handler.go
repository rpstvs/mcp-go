package actions

import (
	"context"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/rpstvs/mcp-go/types"
)

func GetPrAction(copilot *copilot) Handler {
	return func(ctx context.Context, run *types.Run) {
		session, err := copilot.Session(...)
        if err != nil {
            return err
        }

        // perform review

        return nil
	}

}
