package tools

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var AllTools []RegisteredTool

type RegisteredTool struct {
	Tool    mcp.Tool
	Handler server.ToolHandlerFunc
}

func RegisterMCP[I any, O any](name, desc string, handler func(context.Context, mcp.CallToolRequest, *I) (*O, error)) {
	t := mcp.NewTool(name,
		mcp.WithDescription(desc),
		mcp.WithInputSchema[I](),
		mcp.WithOutputSchema[O](),
	)

	h := mcp.NewStructuredToolHandler(handler)

	AllTools = append(AllTools, RegisteredTool{
		Tool:    t,
		Handler: h,
	})
	log.Info("registered tool", "name", name)
}
