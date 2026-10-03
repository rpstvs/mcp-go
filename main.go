package main

import (
	"context"
	"fmt"
	"log"

	copilot "github.com/github/copilot-sdk/go"
)

func main() {
	client := copilot.NewClient(&copilot.ClientOptions{
		LogLevel: "error",
	})

	ctx := context.Background()

	err := client.Start(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Stop()

	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		Model:               "fable",
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer session.Disconnect()

	done := make(chan (bool))

	session.On(func(event copilot.SessionEvent) {
		switch d := event.Data.(type) {
		case *copilot.AssistantMessageData:
			fmt.Println(d.Content)
		case *copilot.SessionIdleData:
			close(done)
		}
	})

	_, err = session.Send(ctx, copilot.MessageOptions{
		Prompt: "cenas",
	})

	if err != nil {
		log.Fatal(err)
	}

	<-done

}
