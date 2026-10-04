package client

import (
	copilot "github.com/github/copilot-sdk/go"
)

type Client struct {
	Client *copilot.Client
}

func newClient() *Client {

	return &Client{
		Client: copilot.NewClient(nil),
	}
}
