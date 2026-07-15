package skills

import (
	"context"
	"embed"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/gbored/golang-mcp-server-template/resources"
)

//go:embed SKILL.md
var SkillFile embed.FS

func init() {
	data, _ := SkillFile.ReadFile("SKILL.md")
	resources.Register("calculator://skills", "Calculator Skill", "Calculator tool usage guidelines",
		func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      req.Params.URI,
					MIMEType: "text/markdown",
					Text:     string(data),
				},
			}, nil
		},
	)
}
