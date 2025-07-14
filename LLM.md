# LLML - Lightweight Language Markup Language API Reference

**The compositional primitive for AI context engineering**

LLML transforms nested data structures into structured markup optimized for AI understanding and performance. Available in Python, TypeScript/JavaScript, Rust, and Go.

## Core Concept

LLML converts data structures to XML-like markup using a formatter system. The default "VibeXML" format is optimized for LLM attention with clear tag boundaries and numbered list items.

## Python API

### Installation
```bash
pip install zenbase-llml
```

### Main Function
```python
from zenbase_llml import llml

def llml(data: Any = None, formatters: Formatters = vibe_xml) -> str
```
- **Purpose**: Core LLML function that formats data using a formatter map
- **Parameters**: 
  - `data`: Any Python object to format
  - `formatters`: Formatter map (defaults to `vibe_xml`)
- **Returns**: Formatted string representation

### Formatters
```python
from zenbase_llml.formatters import vibe_xml, json

# VibeXML formatter (default) - XML-like structured output
vibe_xml: Formatters

# JSON formatter - standard JSON output  
json: Formatters
```

### Type Definitions
```python
from zenbase_llml.formatters import Predicate, FormatFunction, Formatters

Predicate = Callable[[Any], bool]
FormatFunction = Callable[[Any, Callable, Optional[Formatters]], str]  
Formatters = dict[Predicate, FormatFunction]
```

### Usage Examples
```python
from zenbase_llml import llml
from zenbase_llml.formatters import json

# Default VibeXML formatting
result = llml({"name": "John", "items": [1, 2, 3]})
# Output: <name>John</name>\n<items>\n  <items-1>1</items-1>\n  <items-2>2</items-2>\n  <items-3>3</items-3>\n</items>

# JSON formatting  
result = llml(data, json)

# No arguments returns empty string
result = llml()  # ""
```

## TypeScript/JavaScript API

### Installation
```bash
npm install @zenbase/llml
```

### Main Function
```typescript
import { llml } from "@zenbase/llml"

function llml(data: unknown, formatters?: Formatters): string
```
- Converts any data structure to XML-like markup
- Uses VibeXML formatters by default if none provided
- Returns empty string for undefined data

### Type Definitions
```typescript
// Predicate function for type checking
type Predicate = (value: unknown) => boolean

// Formatter function signature
type Formatter = (
  value: unknown,
  llml: (data: unknown, formatters: Formatters) => string,
  formatters: Formatters,
) => string

// Collection of formatter pairs
type Formatters = Iterable<[Predicate, Formatter]>

// Configuration options
interface VibeXMLOptions {
  indent?: string
  prefix?: string
  formatters?: Formatters
}
```

### Built-in Formatters
```typescript
import { formatters } from "@zenbase/llml"

// JSON formatter
formatters.json(
  replacer?: (number | string)[] | ((key: string, value: any) => any) | null,
  space?: string | number
): Formatters

// VibeXML formatter (default)
formatters.vibeXML(options?: VibeXMLOptions): Formatters
```

### Usage Examples
```typescript
import { llml, formatters } from "@zenbase/llml"

// Basic usage
const result = llml({ name: "John", age: 30 })

// JSON output
const jsonResult = llml(data, formatters.json(null, 2))

// Custom formatters
const customFormatters = formatters.vibeXML({
  formatters: new Map([
    [(v) => typeof v === "string", (v) => `<text>${v}</text>`]
  ])
})
const customResult = llml(data, customFormatters)
```

## Rust API

### Installation
```toml
[dependencies]
zenbase-llml = "0.2.0"
```

### Main Functions
```rust
use zenbase_llml::{llml, llml_with_options, LLMLOptions, Prompt};

// Main function
fn llml<T: Prompt>(data: &T) -> String

// With options
fn llml_with_options(data: &Value, options: Option<LLMLOptions>) -> String
```

### Core Trait
```rust
pub trait Prompt {
    fn to_prompt(&self) -> String;
}
```

### Configuration
```rust
#[derive(Debug, Clone, Default)]
pub struct LLMLOptions {
    pub indent: String,    // Indentation string for nested elements
    pub prefix: String,    // Prefix to prepend to all tags
    pub strict: bool,      // Include parent keys as prefixes in nested objects
}
```

### Usage Examples
```rust
use zenbase_llml::llml;
use serde_json::json;

// Basic usage
let data = json!({"key": "value"});
println!("{}", llml(&data));  // "<key>value</key>"

// With options
use zenbase_llml::{llml_with_options, LLMLOptions};

let options = Some(LLMLOptions {
    indent: "  ".to_string(),
    prefix: "app".to_string(),
    strict: false,
});
let result = llml_with_options(&data, options);

// Custom types
struct User { name: String }

impl Prompt for User {
    fn to_prompt(&self) -> String {
        format!("<user>{}</user>", self.name)
    }
}
```

## Go API

### Installation
```bash
go get github.com/zenbase-ai/llml/go
```

### Import
```go
import "github.com/zenbase-ai/llml/go/pkg/llml"
```

### Main Function
```go
func Sprintf(data interface{}, opts ...Options) string
```

### Configuration
```go
type Options struct {
    Indent string  // Indentation string for nested elements
    Prefix string  // Prefix for key names
    Strict bool    // Enable strict mode for nested key prefixes
}
```

### Usage Examples
```go
// Basic usage
result := llml.Sprintf(map[string]any{
    "name": "John",
    "age": 30,
})
// Output: <name>John</name>\n<age>30</age>

// With options
result := llml.Sprintf(data, llml.Options{
    Indent: "  ",
    Prefix: "app",
    Strict: false,
})

// Arrays
result := llml.Sprintf(map[string]any{
    "items": []any{"a", "b", "c"},
})
// Output: <items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n  <items-3>c</items-3>\n</items>
```

## Common Patterns

### VibeXML Format Rules
- **Objects**: `{"key": "value"}` → `<key>value</key>`
- **Arrays**: `{"items": ["a", "b"]}` → `<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n</items>`
- **Nested Objects**: Recursive processing with proper indentation
- **Empty Values**: `{}` and `[]` return empty strings
- **Key Formatting**: Converts to kebab-case (`userAge` → `user-age`)

### Best Practices
1. **Composition**: Build complex prompts from simple, reusable data structures
2. **Type Safety**: Use language-specific type systems for better maintainability
3. **Formatting**: Choose VibeXML for LLM optimization, JSON for API compatibility
4. **Extensibility**: Create custom formatters for specialized use cases

### Performance Notes
- Minimal overhead compared to string concatenation
- Optimized for both human readability and AI model attention
- Structured output reduces prompt ambiguity and improves AI performance

## Migration Guide

### From String Concatenation
```python
# Before
prompt = "Role: " + role + "\n" + "Context: " + json.dumps(context)

# After
prompt = llml({"role": role, "context": context})
```

### From Template Engines
```typescript
// Before (Handlebars/Jinja2)
const template = "{{#each items}}<item>{{this}}</item>{{/each}}"

// After
const result = llml({items: ["a", "b", "c"]})
```

LLML provides a unified, type-safe approach to structured prompt generation across multiple programming languages, optimized for AI context engineering.
