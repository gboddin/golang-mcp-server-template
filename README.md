# golang-mcp-server-template

A minimal MCP server template that takes 2 numbers as input and outputs their addition and subtraction results in separate fields.

## What It Does

Exposes a single MCP tool `add_subtract` that accepts two numbers (`a` and `b`) and returns:
- `addition` — the sum of the two numbers
- `subtraction` — the difference between the two numbers

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `PORT` | No | HTTP listen port (default: `8080`) |

## Usage

```bash
PORT=8080 ./golang-mcp-server-template
```

## System Prompt

The server exposes a `calculator://skills` resource containing usage guidelines. Load this resource before using the tool.

### Note

I'm using this as a base to write various of my MCP servers, it meets my requirements including:

- Minimal boilerplate code, just detail struct fields, write tools and register them
- Can switch between http/stdio/sse for free
- Can embed resources
- Has middleware support for various http auth
- Typed and declarative language, easy to understand for AI
- Create something useful under 5 minutes 

Honestly github.com/mark3labs/mcp-go, tools/all.go and resources/resources.go do most of the work.
