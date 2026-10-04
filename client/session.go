package client

import (
	"context"

	copilot "github.com/github/copilot-sdk/go"
)

func (c *Client) newSession() (*copilot.Session, error) {
	session, err := c.Client.CreateSession(context.Background(), &copilot.SessionConfig{})

	if err != nil {
		return nil, err
	}
	return session, nil
}
