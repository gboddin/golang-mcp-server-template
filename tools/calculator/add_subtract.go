package calculator

import (
	"context"

	"github.com/gbored/golang-mcp-server-template/tools"
	"github.com/mark3labs/mcp-go/mcp"
)

type AddSubstractInput struct {
	A float64 `json:"a" jsonschema:"required,first number"`
	B float64 `json:"b" jsonschema:"required,second number"`
}

type AddSubstractOutput struct {
	Addition     float64 `json:"addition" jsonschema:"result of the addition"`
	Substraction float64 `json:"substraction" jsonschema:"results of the substraction"`
}

func AddSubstractHandler(ctx context.Context, req mcp.CallToolRequest, input *AddSubstractInput) (*AddSubstractOutput, error) {
	return &AddSubstractOutput{
		Addition:     input.A + input.B,
		Substraction: input.A - input.B,
	}, nil
}

func init() {
	tools.RegisterMCP(
		"add_subtract", "Takes two numbers and returns their addition and substraction results", AddSubstractHandler,
	)
}
