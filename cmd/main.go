package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rpstvs/mcp-go/config"
	agentdiscovery "github.com/rpstvs/mcp-go/internal/agent-discovery"
)

func main() {

	config := config.Load()

	app := App{
		config: *config,
	}

	log.Println("MAIN STARTED")
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-example",
		Version: "1.0.0",
	}, nil)

	log.Println("TOOL REGISTERED")
	log.Println("MCP server starting...")

	err := server.Run(context.Background(), &mcp.StdioTransport{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("MCP SERVER IS RUNNING")
}

type App struct {
	config config.Config
	Agents map[string]agentdiscovery.Agent
}
