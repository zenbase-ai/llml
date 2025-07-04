# LLML Trait System Documentation

## Overview

LLML Rust uses a trait-based system that makes it incredibly extensible while maintaining type safety and performance. The core `Prompt` trait allows any type to define how it should be formatted for AI prompts and contexts.

## The `Prompt` Trait

### Definition

```rust
pub trait Prompt {
    fn to_prompt(&self) -> String;
}
```

The `Prompt` trait is intentionally simple - it has just one method that converts a value to its prompt representation.

### Philosophy

The trait-based approach follows Rust's zero-cost abstraction principle:
- **Compile-time dispatch**: No runtime overhead for trait method calls
- **Type safety**: Catch formatting errors at compile time
- **Extensibility**: Any type can participate in the LLML ecosystem
- **Composability**: Build complex prompts from simple, reusable components

## Built-in Implementations

LLML provides `Prompt` implementations for common types:

### Primitive Types

```rust
impl Prompt for String { /* uses to_string() */ }
impl Prompt for &str { /* uses to_string() */ }
impl Prompt for i32 { /* uses to_string() */ }
impl Prompt for i64 { /* uses to_string() */ }
impl Prompt for f32 { /* uses to_string() */ }
impl Prompt for f64 { /* uses to_string() */ }
impl Prompt for bool { /* uses to_string() */ }
```

### JSON Values

```rust
impl Prompt for serde_json::Value {
    fn to_prompt(&self) -> String {
        // Uses full LLML formatting with XML-like tags,
        // kebab-case conversion, and array handling
    }
}
```

## Custom Implementations

### Basic Custom Type

```rust
use zenbase_llml::{llml, Prompt};

#[derive(Debug)]
struct User {
    id: u32,
    name: String,
    email: String,
}

impl Prompt for User {
    fn to_prompt(&self) -> String {
        format!(
            "<user id=\"{}\">\n  <name>{}</name>\n  <email>{}</email>\n</user>",
            self.id, self.name, self.email
        )
    }
}

let user = User {
    id: 123,
    name: "Alice".to_string(),
    email: "alice@example.com".to_string(),
};

println!("{}", llml(&user));
// Output:
// <user id="123">
//   <name>Alice</name>
//   <email>alice@example.com</email>
// </user>
```

### Conditional Formatting

```rust
#[derive(Debug)]
struct TaskStatus {
    task: String,
    completed: bool,
    priority: Priority,
}

#[derive(Debug)]
enum Priority {
    Low,
    Medium,
    High,
    Critical,
}

impl Prompt for Priority {
    fn to_prompt(&self) -> String {
        match self {
            Priority::Low => "🟢 low",
            Priority::Medium => "🟡 medium", 
            Priority::High => "🟠 high",
            Priority::Critical => "🔴 critical",
        }.to_string()
    }
}

impl Prompt for TaskStatus {
    fn to_prompt(&self) -> String {
        let status_icon = if self.completed { "✅" } else { "⏳" };
        let priority = llml(&self.priority);
        
        format!(
            "<task status=\"{}\">\n  <description>{}</description>\n  <priority>{}</priority>\n</task>",
            status_icon, self.task, priority
        )
    }
}
```

### Collections and Complex Types

```rust
use std::collections::HashMap;

#[derive(Debug)]
struct ProjectConfig {
    name: String,
    version: String,
    dependencies: HashMap<String, String>,
    features: Vec<String>,
}

impl Prompt for ProjectConfig {
    fn to_prompt(&self) -> String {
        let mut result = format!(
            "<project>\n  <name>{}</name>\n  <version>{}</version>\n",
            self.name, self.version
        );
        
        // Handle dependencies
        if !self.dependencies.is_empty() {
            result.push_str("  <dependencies>\n");
            for (name, version) in &self.dependencies {
                result.push_str(&format!("    <dependency name=\"{}\">{}</dependency>\n", name, version));
            }
            result.push_str("  </dependencies>\n");
        }
        
        // Handle features
        if !self.features.is_empty() {
            result.push_str("  <features>\n");
            for (i, feature) in self.features.iter().enumerate() {
                result.push_str(&format!("    <feature-{}>{}</feature-{}>\n", i+1, feature, i+1));
            }
            result.push_str("  </features>\n");
        }
        
        result.push_str("</project>");
        result
    }
}

let config = ProjectConfig {
    name: "my-app".to_string(),
    version: "1.2.3".to_string(),
    dependencies: {
        let mut deps = HashMap::new();
        deps.insert("serde".to_string(), "1.0".to_string());
        deps.insert("tokio".to_string(), "1.0".to_string());
        deps
    },
    features: vec!["async".to_string(), "json".to_string()],
};

println!("{}", llml(&config));
```

## Advanced Patterns

### Composition with Other Types

```rust
struct AIContext {
    system_prompt: String,
    user_query: String,
    retrieved_docs: Vec<Document>,
    constraints: Vec<String>,
}

#[derive(Debug)]
struct Document {
    title: String,
    content: String,
    relevance: f64,
}

impl Prompt for Document {
    fn to_prompt(&self) -> String {
        format!(
            "<document relevance=\"{:.2}\">\n  <title>{}</title>\n  <content>{}</content>\n</document>",
            self.relevance, self.title, self.content
        )
    }
}

impl Prompt for AIContext {
    fn to_prompt(&self) -> String {
        let mut result = format!("<context>\n  <system>{}</system>\n", self.system_prompt);
        
        // Use llml() to format documents - this composes with Document's Prompt impl
        if !self.retrieved_docs.is_empty() {
            result.push_str("  <knowledge>\n");
            for doc in &self.retrieved_docs {
                let doc_formatted = llml(doc);
                // Indent the document content
                let indented = doc_formatted.lines()
                    .map(|line| format!("    {}", line))
                    .collect::<Vec<_>>()
                    .join("\n");
                result.push_str(&format!("{}\n", indented));
            }
            result.push_str("  </knowledge>\n");
        }
        
        if !self.constraints.is_empty() {
            result.push_str("  <constraints>\n");
            for (i, constraint) in self.constraints.iter().enumerate() {
                result.push_str(&format!("    <constraint-{}>{}</constraint-{}>\n", i+1, constraint, i+1));
            }
            result.push_str("  </constraints>\n");
        }
        
        result.push_str(&format!("  <query>{}</query>\n</context>", self.user_query));
        result
    }
}
```

### Generic Implementations

```rust
// Generic implementation for any type that implements Debug and Clone
struct DebugWrapper<T: std::fmt::Debug + Clone> {
    value: T,
    label: String,
}

impl<T: std::fmt::Debug + Clone> Prompt for DebugWrapper<T> {
    fn to_prompt(&self) -> String {
        format!("<debug label=\"{}\">{:?}</debug>", self.label, self.value)
    }
}

// Generic implementation for Results
impl<T: Prompt, E: std::fmt::Display> Prompt for Result<T, E> {
    fn to_prompt(&self) -> String {
        match self {
            Ok(value) => format!("<success>{}</success>", llml(value)),
            Err(error) => format!("<error>{}</error>", error),
        }
    }
}

// Generic implementation for Options
impl<T: Prompt> Prompt for Option<T> {
    fn to_prompt(&self) -> String {
        match self {
            Some(value) => llml(value),
            None => "<none/>".to_string(),
        }
    }
}
```

## Best Practices

### When to Use `Prompt` vs `llml_with_options`

**Use `Prompt` trait when:**
- You have custom types that need specific formatting
- You want compile-time guarantees about formatting
- You're building reusable components
- You want to compose types together
- Performance is critical (zero-cost abstractions)

**Use `llml_with_options` when:**
- Working with `serde_json::Value` data
- You need custom indentation, prefixes, or strict mode
- Converting existing JSON-based code
- Prototyping or one-off formatting needs

### Naming Conventions

```rust
// Good: Descriptive, clear purpose
impl Prompt for DatabaseConnection { ... }
impl Prompt for UserPermissions { ... }
impl Prompt for APIResponse { ... }

// Avoid: Generic or unclear names
impl Prompt for Data { ... }
impl Prompt for Thing { ... }
impl Prompt for Wrapper { ... }
```

### Error Handling

```rust
// Handle potential errors gracefully
impl Prompt for UserAccount {
    fn to_prompt(&self) -> String {
        let status = match &self.status {
            AccountStatus::Active => "active",
            AccountStatus::Suspended { reason } => &format!("suspended: {}", reason),
            AccountStatus::Deleted => "deleted",
        };
        
        format!("<account status=\"{}\">{}</account>", status, self.username)
    }
}
```

### Performance Considerations

```rust
// Efficient: Use String::with_capacity for known sizes
impl Prompt for LargeDataSet {
    fn to_prompt(&self) -> String {
        let mut result = String::with_capacity(1024); // Estimate size
        result.push_str("<dataset>\n");
        for item in &self.items {
            result.push_str(&format!("  <item>{}</item>\n", item));
        }
        result.push_str("</dataset>");
        result
    }
}

// Efficient: Reuse allocations when possible
impl Prompt for MetricCollection {
    fn to_prompt(&self) -> String {
        self.metrics.iter()
            .map(|m| format!("<metric name=\"{}\">{}</metric>", m.name, m.value))
            .collect::<Vec<_>>()
            .join("\n")
    }
}
```

## Migration Guide

### From Pattern Matching to Traits

**Before (pattern matching approach):**
```rust
fn format_custom_type(data: &CustomType) -> String {
    match data.type_field {
        TypeA => format!("Type A: {}", data.value),
        TypeB => format!("Type B: {}", data.value),
        TypeC => format!("Type C: {}", data.value),
    }
}
```

**After (trait-based approach):**
```rust
impl Prompt for CustomType {
    fn to_prompt(&self) -> String {
        match self.type_field {
            TypeA => format!("Type A: {}", self.value),
            TypeB => format!("Type B: {}", self.value), 
            TypeC => format!("Type C: {}", self.value),
        }
    }
}

// Now just use:
let result = llml(&my_custom_type);
```

### Leveraging Composition

**Before (monolithic formatting):**
```rust
fn format_user_with_posts(user: &User, posts: &[Post]) -> String {
    let mut result = format!("<user>{}</user>\n", user.name);
    result.push_str("<posts>\n");
    for post in posts {
        result.push_str(&format!("  <post>{}</post>\n", post.title));
    }
    result.push_str("</posts>");
    result
}
```

**After (compositional traits):**
```rust
impl Prompt for User { ... }
impl Prompt for Post { ... }

struct UserWithPosts {
    user: User,
    posts: Vec<Post>,
}

impl Prompt for UserWithPosts {
    fn to_prompt(&self) -> String {
        let user_prompt = llml(&self.user);
        let posts_prompt = self.posts.iter()
            .map(|post| format!("  {}", llml(post)))
            .collect::<Vec<_>>()
            .join("\n");
        
        format!("{}\n<posts>\n{}\n</posts>", user_prompt, posts_prompt)
    }
}
```

## Testing Your Implementations

```rust
#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_user_prompt_format() {
        let user = User {
            id: 123,
            name: "Alice".to_string(),
            email: "alice@example.com".to_string(),
        };
        
        let result = llml(&user);
        
        assert!(result.contains("<user id=\"123\">"));
        assert!(result.contains("<name>Alice</name>"));
        assert!(result.contains("<email>alice@example.com</email>"));
        assert!(result.contains("</user>"));
    }
    
    #[test]
    fn test_empty_collections() {
        let config = ProjectConfig {
            name: "test".to_string(),
            version: "1.0.0".to_string(),
            dependencies: HashMap::new(),
            features: Vec::new(),
        };
        
        let result = llml(&config);
        
        // Should not contain empty dependency or feature sections
        assert!(!result.contains("<dependencies>"));
        assert!(!result.contains("<features>"));
    }
}
```

## Further Reading

- [Main README](../README.md) - Getting started with LLML
- [Rust Trait Documentation](https://doc.rust-lang.org/book/ch10-02-traits.html) - Understanding Rust traits
- [Zero-Cost Abstractions](https://blog.rust-lang.org/2015/05/11/traits.html) - Rust's approach to performance