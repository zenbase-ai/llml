# LLML Project - Claude AI Documentation

## Core LLML Specification

Read .cursor/rules/spec.mdc

## Coding Rules

Reference .cursor/rules/{language = ["py", "ts", "go", "rs"]}.mdc

## Project Overview

LLML (Lightweight Language Markup Language) is a multi-language data serialization library that converts nested data structures into human-readable, XML-like markup. The project provides identical functionality across four programming languages: Python, TypeScript/JavaScript, Rust, and Go.

### Core Purpose
- **AI Prompt Engineering**: Structure complex, multi-part prompts for Large Language Models
- **Configuration Generation**: Create clean, readable configuration files from data
- **Structured Document Creation**: Generate documents with clear hierarchy from data structures
- **Data Serialization**: Convert data into more readable format than JSON/YAML for specific applications

### Key Benefits
- **Zero Configuration**: Works out-of-the-box with sensible defaults
- **Consistent Output**: Identical results across all language implementations
- **Human-Readable**: Produces clean, well-formatted markup
- **LLM-Friendly**: Optimized for AI model consumption and understanding

## Project Structure

```
/Users/knrz/Git/zenbase-ai/llml/
├── README.md                 # Main project documentation
├── justfile                  # Cross-language task runner
├── py/                       # Python implementation
│   ├── README.md
│   ├── pyproject.toml
│   ├── src/
│   │   ├── llml.py          # Main implementation
│   │   └── utils.py         # Helper utilities
│   └── tests/
├── ts/                       # TypeScript implementation
│   ├── README.md
│   ├── package.json
│   ├── src/
│   │   ├── index.ts         # Main implementation
│   │   └── utils.ts         # Helper utilities
│   └── tests/
├── rs/                       # Rust implementation
│   ├── README.md
│   ├── Cargo.toml
│   ├── src/
│   │   ├── lib.rs           # Main implementation
│   │   └── formatters.rs    # Core formatting logic
│   └── tests/
└── go/                       # Go implementation
    ├── README.md
    ├── go.mod
    ├── pkg/llml/
    │   └── llml.go          # Main implementation
    └── tests/
```

Each language implementation is completely self-contained with its own:
- Source code and utilities
- Dependency management files
- Comprehensive test suites
- Language-specific documentation

## Development Environment

### Task Runner: Just
The project uses `just` as a cross-platform task runner for coordinating commands across all implementations.

#### Key Commands
```bash
# Run tests for all languages
just test

# Language-specific commands
just py <command>    # Run command in Python environment (uv)
just ts <command>    # Run command in TypeScript environment (bun)
just go <command>    # Run command in Go environment (go)
just rs <command>    # Run command in Rust environment (cargo)

# AI commands
just claude
just claude -p "your prompt here"      # Run Claude Code CLI
just gemini
just gemini -a -p "your prompt here"   # Run Gemini CLI
```

### CLI Tools

- Instead of `find`, use `fd`
- Instead of `grep`, use `rg`
- Instead of `sed`, use `rg+(awk|sed)` or ast-grep: `ast-grep -p '$A && $A()' -r '$A?.()'`

### Development Tools by Language
- **Python**: `uv` for dependency management, `pytest` for testing, `ruff` for linting
- **TypeScript**: `bun` for runtime and package management, `vitest` for testing
- **Rust**: `cargo` for dependency management and testing
- **Go**: Standard Go toolchain (`go test`, `go mod`)

### 7. Configuration Options
- **Indentation**: Custom indentation strings for nested elements
- **Prefix**: Namespace all generated tags with a prefix

## Testing & Quality Assurance

### Cross-Language Consistency
The test suites ensure that all four implementations produce identical output for the same input data, maintaining consistency across the entire project.

## Usage Examples

### 1. AI Prompt Engineering
```python
# Python
prompt = llml(
    role="Senior Developer",
    task="Code review the following function",
    criteria=["Performance", "Readability", "Best practices"],
    context={"language": "Python", "framework": "FastAPI"}
)
```

### 2. Configuration Generation
```typescript
// TypeScript
const config = llml({
    database: { host: "localhost", port: 5432 },
    features: ["logging", "caching", "monitoring"],
    environment: "production"
});
```

### 3. RAG Chatbot Context
```rust
// Rust
let rag_context = llml(&json!({
    "system": "You are a helpful documentation assistant",
    "documents": [
        {"title": "API Guide", "content": "...", "relevance": 0.95},
        {"title": "Rate Limits", "content": "...", "relevance": 0.82}
    ],
    "query": "How do I authenticate?"
}));
```

### 4. Complex Workflow Definition
```go
// Go
agentPrompt := llml.Sprintf(map[string]any{
    "role": "DevOps Agent",
    "workflows": map[string]any{
        "deploy": []any{
            "Run health checks",
            "Create backup",
            "Deploy to canary",
            "Monitor metrics"
        }
    }
})
```

## Best Practices for Development

1. **Testing**: Always run `just test` or `just test {rs|py|ts|go}` before committing changes
2. **Consistency**: Ensure all language implementations produce identical output
3. **Documentation**: Update README.md, CLAUDE.md, and language-specific README.mds when adding features
4. **Type Safety**: Leverage each language's type system for better reliability
5. **Performance**: Consider memory efficiency for large data structures

## Future Considerations

- **Streaming Support**: For very large data structures
- **Custom Formatters**: Plugin system for specialized formatting
- **Schema Validation**: Optional validation against predefined schemas
- **Output Formats**: Additional output formats beyond XML-like markup
