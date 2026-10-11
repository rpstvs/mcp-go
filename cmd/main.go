package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rpstvs/mcp-go/config"
	"github.com/rpstvs/mcp-go/internal/actions"
	"github.com/rpstvs/mcp-go/internal/copilot"
	"github.com/rpstvs/mcp-go/internal/engine"
	mcpserver "github.com/rpstvs/mcp-go/internal/mcp"
	"github.com/rpstvs/mcp-go/internal/store"
	"github.com/rpstvs/mcp-go/internal/workflow"
)

func main() {
	//Load config
	config := config.Load()
	//Create store
	store := store.NewInMemStore()
	//Create Copilot Client
	copilotFactory := copilot.NewCopilotSessionFactory()

	copilotFactory.Start(context.Background())

	defer copilotFactory.Stop()
	//Create Actions
	actions := actions.NewActions()

	//create workflows
	Workflows := workflow.CreateWorkflows(actions)

	//Create Engine and register workflows
	engine := engine.NewEngine(*config, store, copilotFactory)

	engine.RegisterWorkflows(Workflows)

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
