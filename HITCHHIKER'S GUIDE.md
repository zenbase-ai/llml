# The LLML Guide
*A Practical Introduction to Lightweight Language Markup Language*

> "The secret to good data formats is making complex structures simple to understand."

## Chapter 1: What Is LLML?

### The Problem with Current Formats

Modern applications generate complex, nested data structures that need to be shared between humans and AI systems. JSON works well for machine-to-machine communication but can be hard to read and understand at scale. YAML is more human-readable but can become unwieldy with deep nesting. XML is explicit but verbose and complex.

LLML (Lightweight Language Markup Language) is a data serialization format designed to bridge this gap. It transforms nested data structures into clean, readable markup that's optimized for both human comprehension and AI processing.

### What LLML Does

LLML converts your data structures into XML-like markup using a **formatter system**. Instead of hard-coded rules, LLML uses customizable formatters that determine how different data types are rendered.

Here's the basic API:

```typescript
import { llml } from "@zenbase/llml"

// Basic usage with default formatters
llml({ message: "Hello, World!" })
// Returns: <message>Hello, World!</message>

// Custom formatters (advanced usage)
llml(data, customFormatters)
```

The key insight is that LLML uses **formatters** - pairs of predicate functions and format functions - to determine how each piece of data should be rendered. This makes LLML incredibly flexible while maintaining clean, consistent output.

### Key Features

- **Self-documenting**: Each piece of data is wrapped in descriptive tags
- **Formatter-based**: Extensible system for custom data type handling
- **Clean output**: No unnecessary complexity or verbose syntax
- **AI-optimized**: Structured for clear boundaries and semantic understanding
- **Class-aware**: Automatically uses `.toString()` methods on class instances

## Chapter 2: How LLML Works

### The Formatter System

LLML doesn't use fixed rules for transformation. Instead, it uses a **formatter system** where each formatter consists of:

1. **Predicate function**: Tests if a value matches a specific type
2. **Format function**: Converts the value to string representation

When you call `llml(data)`, it:
1. Iterates through formatters in order
2. Finds the first predicate that returns `true` for your data
3. Uses that formatter's format function to render the value

### Default SwagXML Formatters

The default `swagXML()` formatters handle common data types:

```typescript
import { llml, swagXML } from "@zenbase/llml"

// Get the default formatters
const defaultFormatters = swagXML()
console.log(defaultFormatters) // Map with built-in formatters

// Use explicitly (same as calling llml(data))
llml(data, defaultFormatters)
```

The default formatters handle:
- **Strings**: Plain text content
- **Numbers**: Numeric values
- **Booleans**: `true`/`false`
- **Null**: `"null"`
- **Undefined**: `""`
- **Dates**: Date objects using `.toString()`
- **Arrays**: Numbered item sequences
- **Objects**: Nested tag structures
- **Classes**: Uses `.toString()` method if available

### Basic Transformation Examples

**Simple values:**
```typescript
llml({ user: "Alice", age: 30, active: true })
```
```xml
<user>Alice</user>
<age>30</age>
<active>true</active>
```

**Arrays become numbered sequences:**
```typescript
llml({ tasks: ["code", "test", "deploy"] })
```
```xml
<tasks>
  <tasks-1>code</tasks-1>
  <tasks-2>test</tasks-2>
  <tasks-3>deploy</tasks-3>
</tasks>
```

**Nested objects create hierarchies:**
```typescript
llml({
  user: {
    profile: { name: "Alice", role: "developer" },
    settings: { theme: "dark", notifications: true }
  }
})
```
```xml
<user>
  <profile>
    <name>Alice</name>
    <role>developer</role>
  </profile>
  <settings>
    <theme>dark</theme>
    <notifications>true</notifications>
  </settings>
</user>
```

### Class Instance Handling

LLML automatically uses `.toString()` methods on class instances:

```typescript
class User {
  constructor(public name: string, public email: string) {}

  toString() {
    return `${this.name} (${this.email})`
  }
}

llml({ admin: new User("Alice", "alice@example.com") })
// Output: <admin>Alice (alice@example.com)</admin>
```

### Empty Value Handling

Empty arrays and objects produce no output:

```typescript
llml({ items: [], config: {}, data: undefined })
// Returns: ""
```

This keeps output clean and focused on actual data.

## Chapter 3: The Formatter System

### Understanding Formatters

Formatters are the heart of LLML's flexibility. Each formatter is a `[predicate, formatFunction]` pair:

```typescript
import type { Predicate, Formatter } from "@zenbase/llml"

// Example: Custom email formatter
const isEmail: Predicate = (v): v is string => {
  return typeof v === "string" && v.includes("@")
}

const formatEmail: Formatter = (value: string) => `mailto:${value}`

// Use in custom formatter map
import { swagXML } from "@zenbase/llml"
const customFormatters = swagXML({
  formatters: [[isEmail, formatEmail]]
})

llml({ contact: "alice@example.com" }, customFormatters)
// Output: <contact>mailto:alice@example.com</contact>
```

### Creating Custom Formatters

You can extend LLML by adding custom formatters for specific data types:

```typescript
class Temperature {
  constructor(public celsius: number) {}

  get fahrenheit() { return (this.celsius * 9/5) + 32 }

  toString() {
    return `${this.celsius}°C (${this.fahrenheit.toFixed(1)}°F)`
  }
}

// Class instances automatically use toString()
llml({ current: new Temperature(25) })
// Output: <current>25°C (77.0°F)</current>
```

### Formatter Order Matters

Formatters are processed in order - the first matching predicate wins:

```typescript
// More specific formatters should come first
const formatters = swagXML({
  formatters: [
    [isEmail, formatEmail],        // Specific: email strings
    [isString, formatString]       // General: all strings
  ]
})
```

## Chapter 4: Real-World Examples

### AI Prompt Engineering

LLML excels at structuring complex AI prompts from data:

```typescript
import { llml } from "@zenbase/llml"

const aiPrompt = llml({
  role: "Senior Developer",
  task: "Code review the following function",
  criteria: ["Performance", "Readability", "Best practices"],
  context: {
    language: "Python",
    framework: "FastAPI",
    codebaseMaturity: "production"
  },
  previousFeedback: [
    "Focus on error handling",
    "Consider async/await patterns",
    "Check for security vulnerabilities"
  ]
})

// Creates clear sections that help LLMs understand structure
console.log(aiPrompt)
```

### Configuration Management

Transform configuration objects into readable formats:

```typescript
const config = llml({
  database: {
    host: "localhost",
    port: 5432,
    connectionPool: {
      min: 5,
      max: 20,
      idleTimeout: "30s"
    }
  },
  features: ["caching", "logging", "monitoring"],
  environment: "production"
})

// Output creates clear hierarchical structure
// for both human and machine consumption
```

### RAG System Context

Structure complex contextual information for RAG systems:

```typescript
const ragContext = llml({
  system: "You are a helpful documentation assistant",
  userQuery: "How do I authenticate with your API?",
  retrievedDocuments: [
    {
      title: "API Authentication Guide",
      content: "OAuth 2.0 is the industry standard...",
      relevanceScore: 0.95,
      lastUpdated: "2024-01-15"
    }
  ],
  instructions: [
    "Answer based only on provided documents",
    "Cite specific sources when making claims"
  ]
})
```

## Chapter 5: Formatters and Customization

### Built-in SwagXML Options

The default `swagXML()` formatters support several options:

```typescript
import { swagXML } from "@zenbase/llml"

// Basic usage (no options)
llml(data)

// With custom indentation
llml(data, swagXML({ indent: "  " }))

// With namespace prefix
llml(data, swagXML({ prefix: "api" }))

// With custom formatters
llml(data, swagXML({
  formatters: [[isCustomType, formatCustomType]]
}))
```

### Creating Custom Formatters

For specialized data types, create custom formatters:

```typescript
import { swagXML } from "@zenbase/llml"

// Define your data type
interface CurrencyAmount {
  amount: number
  currency: 'USD' | 'EUR' | 'GBP'
}

// Create predicate function
const isCurrency = (v: unknown): v is CurrencyAmount =>
  typeof v === "object" && v !== null &&
  "amount" in v && "currency" in v

// Create format function
const formatCurrency = (value: CurrencyAmount) => {
  const symbols = { USD: '$', EUR: '€', GBP: '£' }
  return `${symbols[value.currency]}${value.amount.toFixed(2)}`
}

// Use with LLML
const formatters = swagXML({
  formatters: [[isCurrency, formatCurrency]]
})

llml({ price: { amount: 29.99, currency: "USD" } }, formatters)
// Output: <price>$29.99</price>
```

### Recursive Formatting

Format functions can call `llml` recursively for complex nested data:

```typescript
const formatProject = (
  value: Project,
  llml: (data: unknown, formatters: unknown) => string,
  formatters: unknown,
) => {
  const projectData = {
    name: value.name,
    tasks: value.tasks,
    deadline: value.deadline
  }
  return llml(projectData, formatters)
}
```

### Advanced Examples

For comprehensive formatter examples, see the [TypeScript Formatters Guide](ts/docs/formatters.md) which covers:

- Type-specific formatters (emails, URLs, phone numbers)
- Conditional formatting based on content
- Security-aware formatters (redacting sensitive data)
- Performance optimizations for large datasets
- Helper functions for formatter management

## Chapter 6: When to Use LLML

### Good Use Cases

Use LLML when you need:

1. **AI model inputs** - LLMs understand structured context better than raw JSON
2. **Human-readable configuration** - When config files need clear structure
3. **Complex prompt generation** - Multi-part AI instructions with clear boundaries
4. **Documentation generation** - Self-explaining data structures
5. **Debugging complex data** - Clear visualization of nested objects
6. **Class instance serialization** - Automatic `.toString()` method usage

### When to Use Other Formats

Don't use LLML when you need:

1. **High-performance parsing** - JSON is faster for machine processing
2. **Compact data transfer** - JSON/binary formats are more bandwidth-efficient
3. **Standard API communication** - Most APIs expect JSON
4. **Simple key-value data** - Basic formats work fine for flat structures
5. **Database storage** - Unless storing prompts or documentation

### The Formatter Advantage

LLML's formatter system provides unique benefits:

- **Extensibility**: Add support for any data type
- **Consistency**: Same formatting rules across your application
- **Maintainability**: Centralized formatting logic
- **Type safety**: TypeScript support for custom formatters

## Chapter 7: Implementation Details

### Multi-Language Support

LLML maintains consistent behavior across four languages:

```python
# Python
from zenbase_llml import llml
result = llml({"name": "Arthur", "age": 42})
```

```typescript
// TypeScript
import { llml } from "@zenbase/llml"
const result = llml({ name: "Arthur", age: 42 })
```

```rust
// Rust
use llml::llml;
let result = llml(&json!({"name": "Arthur", "age": 42}));
```

```go
// Go
import "github.com/zenbase-ai/llml"
result := llml.Sprintf(map[string]any{"name": "Arthur", "age": 42})
```

All implementations produce identical output for the same input data.

### Design Principles

LLML follows these core principles:

1. **Predictability** - Same input always produces same output
2. **Readability** - Data should be self-explanatory
3. **Extensibility** - Formatter system allows customization
4. **Simplicity** - Simple API with powerful capabilities

## Appendix A: Quick Reference

```typescript
import { llml, swagXML } from "@zenbase/llml"

// Basic usage
llml({ message: "Hello" })                    // <message>Hello</message>

// Arrays become numbered sequences
llml({ items: ["a", "b", "c"] })
// <items>
//   <items-1>a</items-1>
//   <items-2>b</items-2>
//   <items-3>c</items-3>
// </items>

// Nested objects show hierarchy
llml({ config: { debug: true, timeout: 30 } })
// <config>
//   <debug>true</debug>
//   <timeout>30</timeout>
// </config>

// Advanced options
llml(data, swagXML({
  indent: "  ",                    // Custom indentation
  prefix: "app",                   // Namespace all tags
  formatters: [[isCustom, formatCustom]]  // Custom formatters
}))

// Empty values are omitted
llml([])              // ""
llml({})              // ""
```

## Appendix B: Installation

**Python**: `pip install zenbase-llml`
```python
from zenbase_llml import llml
result = llml({"key": "value"})
```

**TypeScript/JavaScript**: `npm install @zenbase/llml`
```typescript
import { llml } from "@zenbase/llml"
const result = llml({ key: "value" })
```

**Rust**: `cargo add zenbase_llml`
```rust
use llml::llml;
let result = llml(&json!({"key": "value"}));
```

**Go**: `go get github.com/zenbase-ai/llml/go`
```go
import "github.com/zenbase-ai/llml"
result := llml.Sprintf(map[string]any{"key": "value"})
```

## Appendix C: Learn More

- **[TypeScript Formatters Guide](ts/docs/formatters.md)** - Complete guide to custom formatters
- **[Real Examples](ts/examples/)** - Production use cases and patterns
- **GitHub Repository** - Source code and issue tracking

---

**Version 2.0** - Updated for the modern formatter-based LLML architecture.
