package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/gbored/golang-mcp-server-template/config"
	
	"github.com/gbored/golang-mcp-server-template/resources"
	"github.com/gbored/golang-mcp-server-template/tools"
	_ "github.com/gbored/golang-mcp-server-template/tools/calculator"
)

func main() {
	srv := server.NewMCPServer("calculator-mcp", "1.0.0", server.WithResourceCapabilities(false, false))

	for _, t := range tools.AllTools {
		srv.AddTool(t.Tool, func(ctx context.Context, request mcp.CallToolRequest) (resp *mcp.CallToolResult, err error) {
			log.Debug("calling tool", "name", t.Tool.Name, "args", request.GetArguments())
			resp, err = t.Handler(ctx, request)
			if err != nil {
				log.Warn("failed to call tool", "name", t.Tool.Name, "args", request.GetArguments())
				return
			}
			log.Debug("received tool reply", "name", t.Tool.Name, "resp", resp.StructuredContent)
			return
		})
	}

	for _, r := range resources.AllResources {
		srv.AddResource(r.Resource, r.Handler)
	}

	port := config.Port()
	apiKey := config.APIKey()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != apiKey {
			log.Warn("api key not match", "remote", r.RemoteAddr, "method", r.Method, "path", r.URL.Path)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		log.Info("serving request", "remote", r.RemoteAddr, "method", r.Method, "path", r.URL.Path)
		server.NewStreamableHTTPServer(srv).ServeHTTP(w, r)
	})
	log.Info("Started HTTP server", "port", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func init() {
	if strings.ToLower(os.Getenv("DEBUG")) == "true" ||
		os.Getenv("DEBUG") == "1" ||
		strings.ToLower(os.Getenv("DEBUG")) == "yes" {
		log.SetLevel(log.DebugLevel)
		log.Warn("Running in debug mode, tool call information, including arguments, will be logged")
	}
}
