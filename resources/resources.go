package resources

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var AllResources []RegisteredResource

type RegisteredResource struct {
	Resource mcp.Resource
	Handler  server.ResourceHandlerFunc
}

func Register(uri, name, desc string, handler server.ResourceHandlerFunc) {
	r := mcp.NewResource(uri, name, mcp.WithResourceDescription(desc))
	AllResources = append(AllResources, RegisteredResource{
		Resource: r,
		Handler:  handler,
	})
	log.Info("registered resource", "uri", uri)
}

func RegisterFile(uri, name, desc, filePath string) {
	Register(uri, name, desc, func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      req.Params.URI,
				MIMEType: "text/markdown",
				Text:     string(data),
			},
		}, nil
	})
}
