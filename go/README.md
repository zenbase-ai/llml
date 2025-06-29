# LLML Go

Go implementation of the Lightweight Language Markup Language (LLML) library.

## Overview

LLML is a data serialization library that transforms nested data structures (maps, slices, primitives) into well-formatted, XML-like markup. It's designed for tasks like AI prompt engineering, configuration file generation, and creating structured documents from data.

## Features

- **Automatic kebab-case conversion** of keys
- **Smart list formatting** with numbered items and wrapper tags
- **Recursive handling** of deeply nested structures
- **Automatic indentation** for readability
- **Multiline content support**
- **Optional prefix namespacing**
- **Type-safe Go implementation** with comprehensive type coverage
- **Zero configuration** - works out of the box
- **Strict mode control** - choose whether nested properties include parent key prefixes

## Installation

```bash
go get github.com/zenbase-ai/llml/go
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/zenbase-ai/llml/go/pkg/llml"
)

func main() {
    // Simple example
    result := llml.Sprintf(map[string]any{
        "task": "analyze",
        "content": "customer feedback",
    })
    fmt.Println(result)
    // Output: <task>analyze</task>
    //         <content>customer feedback</content>

    // List handling
    result = llml.Sprintf(map[string]any{
        "rules": []any{"be concise", "be helpful", "be accurate"},
    })
    fmt.Println(result)
    // Output: <rules>
    //           <rules-1>be concise</rules-1>
    //           <rules-2>be helpful</rules-2>
    //           <rules-3>be accurate</rules-3>
    //         </rules>
}
```

## Usage Examples

### Example 1: Data Extraction

```go
extractionPrompt := llml.Sprintf(map[string]any{
    "task":         "Extract key information from customer feedback",
    "instructions": "Identify and categorize customer sentiments and specific issues mentioned",
    "rules": []any{
        "Classify sentiment as positive, negative, or neutral",
        "Extract specific product features mentioned",
        "Identify any requested improvements or fixes",
        "Note any comparisons to competitors",
    },
    "output_format": map[string]any{
        "sentiment":          "positive/negative/neutral",
        "features_mentioned": []any{"list of features"},
        "issues":             []any{"list of problems"},
        "improvements":       []any{"list of suggestions"},
    },
})
fmt.Println(extractionPrompt)
```

Output:
```xml
<instructions>Identify and categorize customer sentiments and specific issues mentioned</instructions>
<output-format>
  <output-format-features-mentioned>
    <features-mentioned-1>list of features</features-mentioned-1>
  </output-format-features-mentioned>
  <output-format-improvements>
    <improvements-1>list of suggestions</improvements-1>
  </output-format-improvements>
  <output-format-issues>
    <issues-1>list of problems</issues-1>
  </output-format-issues>
  <output-format-sentiment>positive/negative/neutral</output-format-sentiment>
</output-format>
<rules>
  <rules-1>Classify sentiment as positive, negative, or neutral</rules-1>
  <rules-2>Extract specific product features mentioned</rules-2>
  <rules-3>Identify any requested improvements or fixes</rules-3>
  <rules-4>Note any comparisons to competitors</rules-4>
</rules>
<task>Extract key information from customer feedback</task>
```

### Example 2: RAG Chatbot

```go
ragPrompt := llml.Sprintf(map[string]any{
    "system":       "You are a helpful documentation assistant",
    "instructions": "Answer questions based on the provided documentation context",
    "documents": []any{
        map[string]any{
            "title":           "API Authentication Guide",
            "content":         "Our API uses OAuth 2.0 for authentication...",
            "relevance_score": 0.95,
        },
        map[string]any{
            "title":           "Rate Limiting Documentation",
            "content":         "API calls are limited to 1000 requests per hour...",
            "relevance_score": 0.82,
        },
    },
    "user_query": "How do I authenticate with your API?",
    "constraints": []any{
        "Only use information from the provided documents",
        "Cite the document title when referencing information",
        "If information is not available, explicitly state so",
    },
})
fmt.Println(ragPrompt)
```

### Example 3: AI Agent with Workflows

```go
agentPrompt := llml.Sprintf(map[string]any{
    "role": "DevOps automation agent",
    "context": map[string]any{
        "environment":     "production",
        "aws_region":      "us-east-1",
        "services":        []any{"web-api", "worker-queue", "database"},
        "last_deployment": "2024-01-15T10:30:00Z",
    },
    "instructions": "Execute deployment workflow with safety checks",
    "workflows": map[string]any{
        "deploy": []any{
            "Run pre-deployment health checks",
            "Create backup of current state",
            "Deploy to canary instance (5% traffic)",
            "Monitor metrics for 10 minutes",
            "If healthy, proceed to full deployment",
            "If issues detected, automatic rollback",
        },
        "rollback": []any{
            "Stop new traffic to affected services",
            "Restore from latest backup",
            "Verify service health",
            "Send notification to ops channel",
        },
    },
    "safety_rules": []any{
        "Never skip health checks",
        "Always maintain 99.9% uptime SLA",
        "Require manual approval for database changes",
    },
})
fmt.Println(agentPrompt)
```

## Configuration Options

The `llml.Sprintf` function accepts optional configuration through the `Options` struct:

```go
type Options struct {
    Indent string  // Indentation string (default: "")
    Prefix string  // Prefix for all tags (default: "")
    Strict bool    // Include parent key prefixes in nested objects (default: false)
}
```

### Using Options

```go
// With custom indentation
result := llml.Sprintf(map[string]any{
    "config": map[string]any{
        "debug":   true,
        "timeout": 30,
    },
}, llml.Options{Indent: "  "})

// With prefix
result := llml.Sprintf(map[string]any{
    "database": "postgres",
    "port":     5432,
}, llml.Options{Prefix: "app"})
// Output: <app-database>postgres</app-database>
//         <app-port>5432</app-port>

// With both
result := llml.Sprintf(map[string]any{
    "items": []any{"first", "second"},
}, llml.Options{
    Indent: "    ",
    Prefix: "config",
})

// Example with strict mode
result := llml.Sprintf(map[string]any{
    "config": map[string]any{
        "debug":   true,
        "timeout": 30,
    },
}, llml.Options{Strict: true})
// Output: <config>
//           <config-debug>true</config-debug>  
//           <config-timeout>30</config-timeout>
//         </config>

// Example with strict mode disabled (default)
result = llml.Sprintf(map[string]any{
    "config": map[string]any{
        "debug":   true,
        "timeout": 30,
    },
}, llml.Options{Strict: false})
// Output: <config>
//           <debug>true</debug>
//           <timeout>30</timeout>
//         </config>
```

## Data Type Support

LLML Go supports all Go data types:

```go
data := map[string]any{
    // Strings
    "message": "Hello, World!",

    // Numbers
    "count":       42,
    "temperature": 98.6,
    "price":       19.99,

    // Booleans
    "enabled":  true,
    "disabled": false,

    // Slices
    "tags":    []any{"go", "llml", "markup"},
    "numbers": []any{1, 2, 3, 4, 5},

    // Nested maps
    "config": map[string]any{
        "debug": true,
        "level": "info",
    },

    // Mixed data
    "users": []any{
        map[string]any{"name": "Alice", "age": 30},
        map[string]any{"name": "Bob", "age": 25},
    },

    // Nil values
    "optional": nil,

    // Empty values
    "empty_string": "",
    "empty_list":   []any{},
    "zero_value":   0,
}

result := llml.Sprintf(data)
```

## Key Formatting

Keys are automatically converted to kebab-case:

```go
result := llml.Sprintf(map[string]any{
    "userName":        "alice",
    "user_email":      "alice@example.com",
    "userPhoneNumber": "555-1234",
    "user age":        30,
})
// Output: <user-age>30</user-age>
//         <user-email>alice@example.com</user-email>
//         <user-name>alice</user-name>
//         <user-phone-number>555-1234</user-phone-number>
```

## Multiline Content

Multiline strings are automatically formatted with proper indentation:

```go
content := `
This is a multiline string
with multiple lines
and proper formatting
`

result := llml.Sprintf(map[string]any{
    "description": content,
})
// Output: <description>
//           This is a multiline string
//           with multiple lines
//           and proper formatting
//         </description>
```

## API Reference

### `llml.Sprintf(data interface{}, opts ...Options) string`

Converts data structures to XML-like markup.

**Parameters:**
- `data`: Input data to convert (maps, slices, primitives, nil)
- `opts`: Optional configuration (indentation, prefix)

**Returns:**
- String containing the formatted XML-like markup

**Behavior:**
- `nil` → `"nil"`
- Empty map → `""`
- Empty slice → `""`
- Maps → Nested tags with kebab-case keys
- Slices → Numbered items with wrapper tags
- Primitives → String representation

### `llml.Options`

Configuration struct for customizing output format.

**Fields:**
- `Indent`: String used for indentation (default: `""`)
- `Prefix`: Prefix added to all tag names (default: `""`)

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Performance

The Go implementation offers several performance benefits:

- **Type Safety**: Compile-time type checking prevents runtime errors
- **Memory Efficiency**: Efficient string building and minimal allocations
- **Concurrent Safe**: No global state, safe for concurrent use
- **Zero Dependencies**: Only uses Go standard library (except testing)

## Development

This project uses:
- Go 1.21+ for language features
- `github.com/stretchr/testify` for testing assertions
- Standard Go testing framework

## Contributing

The Go implementation follows the same behavioral specifications as the Python and TypeScript versions, ensuring consistent output across all language implementations.
