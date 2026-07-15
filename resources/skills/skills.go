package skills

import (
	"context"
	"embed"

	"github.com/gbored/golang-mcp-server-template/resources"
	"github.com/mark3labs/mcp-go/mcp"
)

//go:embed SKILL.md
var SkillFile embed.FS

func init() {
	data, err := SkillFile.ReadFile("SKILL.md")
	if err != nil {
		panic(err)
	}
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
