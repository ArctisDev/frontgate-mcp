package main

import (
	"context"
	"log"
	"os"

	"github.com/ArctisDev/frontgate/internal/mcp"
)

func main() {
	logger := log.New(os.Stderr, "frontgate: ", log.LstdFlags|log.Lmsgprefix)
	logger.Printf("starting stdio MCP server")
	server := mcp.NewServer(logger)
	if err := server.Run(context.Background(), os.Stdin, os.Stdout); err != nil {
		logger.Fatalf("server failed: %v", err)
	}
}
