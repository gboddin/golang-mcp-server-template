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
