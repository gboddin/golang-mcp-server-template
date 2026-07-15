package resources

import (
	"context"
	"embed"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

//go:embed skills/SKILL.md
var skillFile embed.FS

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
		data, err := fs.ReadFile(skillFile, filePath)
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

func FindProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}
