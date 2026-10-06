package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rpstvs/mcp-go/config"
	mcpserver "github.com/rpstvs/mcp-go/internal/mcp"
	"github.com/rpstvs/mcp-go/internal/store"
	"github.com/rpstvs/mcp-go/internal/workflow"
)

func main() {

	config := config.Load()

	store := store.NewInMemStore()

	engine := workflow.NewEngine(*config, store)

	log.Println("MAIN STARTED")
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-example",
		Version: "1.0.0",
	}, nil)

	log.Println("TOOL REGISTERED")
	mcpserver.Register(server, engine)
	log.Println("MCP server starting...")

	err := server.Run(context.Background(), &mcp.StdioTransport{})

	if err != nil {
		log.Fatal(err)
	}
}
