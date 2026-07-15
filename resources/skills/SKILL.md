# Calculator Skill

## Overview

This MCP server provides a simple calculator tool that performs basic arithmetic operations on two numbers.

## Tool: add_subtract

Accepts two numbers and returns their addition and subtraction results.

### Input Schema

| Field | Type | Description |
|-------|------|-------------|
| `a` | number | First number |
| `b` | number | Second number |

### Output Schema

| Field | Type | Description |
|-------|------|-------------|
| `addition` | number | Sum of a and b |
| `subtraction` | number | Difference of a and b |

### Example

```json
{
  "a": 10,
  "b": 3
}
```

Result:
```json
{
  "addition": 13,
  "subtraction": 7
}
```
