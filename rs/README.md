# LLML Rust

Rust implementation of the Lightweight Markup Language (LLML) - **React for Prompts** - a library that converts data structures into XML-like markup with intelligent formatting and extensible trait-based design.

## Features

- **Simple Key-Value Conversion**: `{"key": "value"}` → `<key>value</key>`
- **Array Formatting**: Arrays become numbered lists with wrapper tags
- **Nested Object Support**: Recursive processing with kebab-case key conversion
- **Trait-Based Extensibility**: Implement `Prompt` trait for any custom type
- **Zero Configuration**: Works out-of-the-box with sensible defaults
- **Configurable Indentation**: Customizable spacing and prefixes
- **Multiline String Handling**: Proper formatting for multiline content
- **Kebab-Case Conversion**: Automatic conversion of camelCase and snake_case keys
- **Strict Mode Control**: Choose whether nested properties include parent key prefixes

## Installation

Add this to your `Cargo.toml`:

```toml
[dependencies]
zenbase-llml = "0.2.0"
```

## Quick Start

```rust
use zenbase_llml::llml;
use serde_json::json;

fn main() {
    // Simple key-value pairs
    let data = json!({"instructions": "Follow these steps"});
    println!("{}", llml(&data));
    // Output: <instructions>Follow these steps</instructions>
    
    // Arrays become numbered lists
    let data = json!({"rules": ["first", "second", "third"]});
    println!("{}", llml(&data));
    // Output:
    // <rules>
    //   <rules-1>first</rules-1>
    //   <rules-2>second</rules-2>
    //   <rules-3>third</rules-3>
    // </rules>
    
    // Works with any type!
    println!("{}", llml(&"Hello World"));  // "Hello World"
    println!("{}", llml(&42));             // "42"
    println!("{}", llml(&true));           // "true"
}
```

## Advanced Usage

### Custom Formatting Options

For `serde_json::Value` types, you can use `llml_with_options` for custom formatting:

```rust
use zenbase_llml::{llml_with_options, LLMLOptions};
use serde_json::json;

let data = json!({"message": "Hello World"});
let options = Some(LLMLOptions {
    indent: "  ".to_string(),
    prefix: "app".to_string(),
    strict: false,
});

let result = llml_with_options(&data, options);
// Output:   <app-message>Hello World</app-message>

// Example with strict mode
let data = json!({"config": {"debug": true, "timeout": 30}});
let options = Some(LLMLOptions {
    indent: "".to_string(),
    prefix: "".to_string(),
    strict: true,
});
let result = llml_with_options(&data, options);
// Output: <config>
//           <config-debug>true</config-debug>
//           <config-timeout>30</config-timeout>
//         </config>

// Example with strict mode disabled (default)
let options = Some(LLMLOptions {
    indent: "".to_string(),
    prefix: "".to_string(),
    strict: false,
});
let result = llml_with_options(&data, options);
// Output: <config>
//           <debug>true</debug>
//           <timeout>30</timeout>
//         </config>
```

### Custom Types with the `Prompt` Trait

The real power of LLML Rust comes from implementing the `Prompt` trait for your own types:

```rust
use zenbase_llml::{llml, Prompt};

// Define your custom type
#[derive(Debug)]
struct User {
    name: String,
    role: String,
    active: bool,
}

// Implement the Prompt trait
impl Prompt for User {
    fn to_prompt(&self) -> String {
        format!(
            "<user>\n  <name>{}</name>\n  <role>{}</role>\n  <status>{}</status>\n</user>",
            self.name, 
            self.role, 
            if self.active { "active" } else { "inactive" }
        )
    }
}

fn main() {
    let user = User {
        name: "Alice".to_string(),
        role: "Developer".to_string(),
        active: true,
    };
    
    println!("{}", llml(&user));
    // Output:
    // <user>
    //   <name>Alice</name>
    //   <role>Developer</role>
    //   <status>active</status>
    // </user>
}
```

### AI Agent Configuration Example

```rust
use zenbase_llml::{llml, Prompt};

struct AgentConfig {
    role: String,
    capabilities: Vec<String>,
    constraints: Vec<String>,
}

impl Prompt for AgentConfig {
    fn to_prompt(&self) -> String {
        let mut result = format!("<role>{}</role>\n", self.role);
        
        if !self.capabilities.is_empty() {
            result.push_str("<capabilities>\n");
            for (i, cap) in self.capabilities.iter().enumerate() {
                result.push_str(&format!("  <capability-{}>{}</capability-{}>\n", i+1, cap, i+1));
            }
            result.push_str("</capabilities>\n");
        }
        
        if !self.constraints.is_empty() {
            result.push_str("<constraints>\n");
            for (i, constraint) in self.constraints.iter().enumerate() {
                result.push_str(&format!("  <constraint-{}>{}</constraint-{}>\n", i+1, constraint, i+1));
            }
            result.push_str("</constraints>");
        }
        
        result
    }
}

let agent = AgentConfig {
    role: "Code Review Assistant".to_string(),
    capabilities: vec![
        "Analyze code quality".to_string(),
        "Suggest improvements".to_string(),
        "Check best practices".to_string(),
    ],
    constraints: vec![
        "Focus on constructive feedback".to_string(),
        "Provide specific examples".to_string(),
    ],
};

println!("{}", llml(&agent));
```

### Handling Complex Data Structures

```rust
use zenbase_llml::llml;
use serde_json::json;

let complex_data = json!({
    "user_config": {
        "theme_mode": "dark",
        "max_retries": 5
    },
    "data_sources": [
        {"name": "db1", "active": true},
        {"name": "db2", "active": false}
    ]
});

let result = llml(&complex_data);
// Produces nested XML-like structure with kebab-case keys
```

## Key Transformations

| Input | Output |
|-------|--------|
| `user_name` | `user-name` |
| `maxRetries` | `max-retries` |
| `key with spaces` | `key-with-spaces` |

## Supported Data Types

- **Primitives**: Strings, numbers, booleans, null (via `Display` trait)
- **JSON Values**: Full support via `serde_json::Value` with LLML formatting
- **Arrays**: Converted to numbered lists with wrapper tags
- **Objects**: Recursively processed with nested tag structure
- **Custom Types**: Any type implementing the `Prompt` trait
- **Empty Values**: Handled gracefully (empty objects/arrays → empty strings)

## Examples

Run the included examples to see LLML in action:

```bash
cargo run --example basic_usage
```

This will demonstrate:
- Simple key-value formatting
- Array handling
- Nested objects
- Mixed content
- Custom indentation
- Empty value handling

## Testing

Run the comprehensive test suite:

```bash
cargo test
```

The library includes extensive tests covering:
- Edge cases (empty values, null, zero, false)
- Data type conversions
- Kebab-case transformations
- Nested structures
- Indentation and prefixes
- Multiline content

## Architecture

The library uses a trait-based system with two levels:

### Simple API: `llml<T: Prompt>(data: &T)`
1. **Trait Dispatch**: Uses Rust's trait system to determine formatting
2. **Type Safety**: Compile-time guarantees about supported types
3. **Extensibility**: Any type can implement `Prompt` for custom formatting

### Advanced API: `llml_with_options(&Value, Option<LLMLOptions>)`
1. **JSON-Specific**: Specialized for `serde_json::Value` types
2. **Configurable**: Custom indentation, prefixes, and strict mode
3. **Recursive Processing**: Handles nested structures with intelligent formatting

### Core Features
- **Kebab-case conversion**: Automatic conversion of camelCase and snake_case keys
- **Array handling**: Numbered lists with wrapper tags
- **Multiline support**: Proper formatting for multiline content
- **Empty value handling**: Graceful handling of empty objects/arrays

For detailed information about implementing custom types, see [`docs/traits.md`](docs/traits.md).

## Contributing

This is part of the LLML project. See the main repository for contribution guidelines.

## License

MIT