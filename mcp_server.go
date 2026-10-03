package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func mcpExample() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-example",
		Version: "1.0.0",
	}, nil)

	server.AddTool(&mcp.Tool{
		Name: "greet",
	}, example_handler)

	err := server.Run(context.Background(), &mcp.StdioTransport{})

	if err != nil {
		log.Fatal(err)
	}
}

func example_handler(context.Context, *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return &mcp.CallToolResult{}, nil
}
