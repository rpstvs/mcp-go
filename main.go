package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	log.Println("MAIN STARTED")
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-example",
		Version: "1.0.0",
	}, nil)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "greet_and_add",
			Description: "Greet someone and add two numbers",
		},
		GreetandAddWorkflow,
	)

	log.Println("TOOL REGISTERED")
	log.Println("MCP server starting...")

	err := server.Run(context.Background(), &mcp.StdioTransport{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("MCP SERVER IS RUNNING")
}

func GreetandAddWorkflow(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GreetAndAddInput,
) (*mcp.CallToolResult, GreetAndAddOutput, error) {
	sum := add(input.A, input.B)
	return &mcp.CallToolResult{}, GreetAndAddOutput{Message: "Hello Rui",
		Sum: sum}, nil
}

func add(a, b int) int {
	return a + b
}

type GreetAndAddInput struct {
	Name string `json:"name"`
	A    int    `json:"a"`
	B    int    `json:"b"`
}

type GreetAndAddOutput struct {
	Message string `json:"message"`
	Sum     int    `json:"sum"`
}
