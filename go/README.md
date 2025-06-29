# LLML Go

Go implementation of the Lightweight Markup Language (LLML) library.

## Installation

```bash
go get github.com/zenbase-ai/llml-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/zenbase-ai/llml-go/pkg/llml"
)

func main() {
    data := map[string]interface{}{
        "instructions": "Follow these steps",
        "rules": []string{"first", "second", "third"},
        "config": map[string]interface{}{
            "debug": true,
            "timeout": 30,
        },
    }
    
    result := llml.LLML(data)
    fmt.Println(result)
}
```

## Running Tests

```bash
go test ./...
```