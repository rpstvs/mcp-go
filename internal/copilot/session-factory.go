package copilot

import (
	"context"

	copilot "github.com/github/copilot-sdk/go"
)

type SessionFactory interface {
	Create(ctx context.Context) (*copilot.Session, error)
}

type CopilotSessionFactory struct {
	CopilotClient *copilot.Client
}

func NewCopilotSessionFactory() *CopilotSessionFactory {
	client := copilot.NewClient(&copilot.ClientOptions{
		LogLevel: "error",
	})

	return &CopilotSessionFactory{
		CopilotClient: client,
	}
}

func (f *CopilotSessionFactory) Create(ctx context.Context) (*copilot.Session, error) {
	session, err := f.CopilotClient.CreateSession(ctx, &copilot.SessionConfig{})

	if err != nil {
		return nil, err
	}

	return session, nil
}
