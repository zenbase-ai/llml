# 🚀 LLML TypeScript - Lightweight Language Markup Language

**The most elegant way to generate structured text in TypeScript/JavaScript.**

LLML transforms your data into beautifully formatted XML-like markup with zero fuss and maximum flexibility. Perfect for context engineering, configuration generation, and structured document creation.

## ⚡ Quick Start

```typescript
import { llml } from '@zenbase/llml';

// Simple values
console.log(llml({ greeting: "Hello World" }));
// Output: <greeting>Hello World</greeting>

// Lists become numbered items
console.log(llml({ tasks: ["Buy milk", "Walk dog", "Code LLML"] }));
// Output:
// <tasks>
//   <tasks-1>Buy milk</tasks-1>
//   <tasks-2>Walk dog</tasks-2>
//   <tasks-3>Code LLML</tasks-3>
// </tasks>

// Complex nested structures
console.log(llml({
    title: "My Project",
    features: ["Fast", "Simple", "Powerful"],
    config: { debug: true, version: "1.0" }
}));
// Output:
// <title>My Project</title>
// <features>
//   <features-1>Fast</features-1>
//   <features-2>Simple</features-2>
//   <features-3>Powerful</features-3>
// </features>
// <config>
//   <debug>true</debug>
//   <version>1.0</version>
// </config>
```

## 🎯 Why LLML?

- **🔥 Zero Learning Curve**: One function, infinite possibilities
- **🎨 Beautiful Output**: Automatically formatted, properly indented
- **⚡ Lightning Fast**: Minimal overhead, maximum performance
- **🌟 Modern TypeScript**: Fully typed with excellent IDE support
- **🔧 Extensible Formatters**: Customize formatting for any data type
- **🎯 Type-Safe**: Predicate-based formatter system with full type safety

## 🛠️ Installation

```bash
# Using bun (recommended)
bun install

# Using npm
npm install

# Using yarn
yarn install
```

## 📚 Advanced Usage

### Custom Formatters
LLML uses an extensible formatter system. You can create custom formatters for specialized data types:

```typescript
import { llml, vibeXML } from '@zenbase/llml';

// Custom formatters for domain objects
class User {
  constructor(public name: string, public email: string) {}
}

class Money {
  constructor(public amount: number, public currency: string) {}
}

// Create custom formatters
const customFormatters = new Map();
customFormatters.set(
  (v: unknown): v is User => v instanceof User,
  (v: User) => `${v.name} <${v.email}>`
);
customFormatters.set(
  (v: unknown): v is Money => v instanceof Money,
  (v: Money) => `${v.amount} ${v.currency}`
);

// Use with vibeXML
const formatters = vibeXML({ formatters: customFormatters });
const result = llml({
  customer: new User("Alice", "alice@example.com"),
  price: new Money(100, "USD")
}, formatters);
// Output:
// <customer>Alice <alice@example.com></customer>
// <price>100 USD</price>
```

### Formatter Precedence
Custom formatters take precedence over built-in ones. The first matching formatter wins:

```typescript
// Override built-in boolean formatting
const customFormatters = new Map();
customFormatters.set(
  (v: unknown): v is boolean => typeof v === "boolean",
  (v: boolean) => v ? "YES" : "NO"
);

const formatters = vibeXML({ formatters: customFormatters });
const result = llml({ enabled: true, disabled: false }, formatters);
// Output:
// <enabled>YES</enabled>
// <disabled>NO</disabled>
```

### Multi-line Content
```typescript
const instructions = `
Step 1: Install LLML
Step 2: Import llml
Step 3: Create magic
`;

console.log(llml({ instructions }));
// Output:
// <instructions>
// Step 1: Install LLML
// Step 2: Import llml
// Step 3: Create magic
// </instructions>
```

### Complex Nested Structures
```typescript
const promptData = {
    system: "You are a helpful assistant",
    rules: [
        "Be concise and clear",
        "Provide examples when helpful",
        "Ask clarifying questions"
    ],
    context: {
        userLevel: "beginner",
        topic: "TypeScript programming"
    }
};

console.log(llml(promptData));
// Output:
// <system>You are a helpful assistant</system>
// <rules>
//   <rules-1>Be concise and clear</rules-1>
//   <rules-2>Provide examples when helpful</rules-2>
//   <rules-3>Ask clarifying questions</rules-3>
// </rules>
// <context>
//   <userLevel>beginner</userLevel>
//   <topic>TypeScript programming</topic>
// </context>
```

### Built-in Type Support
LLML automatically handles common JavaScript types:

```typescript
// Dates, URLs, and other built-in objects
const data = {
    timestamp: new Date("2023-01-01T00:00:00Z"),
    homepage: new URL("https://example.com"),
    enabled: true,
    count: 42
};

console.log(llml(data));
// Output:
// <timestamp>Sun Jan 01 2023 00:00:00 GMT+0000 (Coordinated Universal Time)</timestamp>
// <homepage>https://example.com/</homepage>
// <enabled>true</enabled>
// <count>42</count>
```

## 🎪 Use Cases

### 🤖 AI Prompt Engineering
Perfect for structuring complex prompts:
```typescript
const prompt = llml({
    role: "Senior TypeScript Developer",
    task: "Code review the following function",
    criteria: ["Performance", "Readability", "Best practices"],
    code: functionToReview
});
```

### ⚙️ Configuration Generation
Generate clean config files:
```typescript
const config = llml({
    database: { host: "localhost", port: 5432 },
    features: ["logging", "caching", "monitoring"],
    environment: "production"
});
```

### 📄 Document Structure
Create structured documents:
```typescript
const document = llml({
    title: "API Documentation",
    sections: ["Authentication", "Endpoints", "Examples"],
    metadata: { version: "2.1", author: "Dev Team" }
});
```

## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run tests once
bun test

# Run tests in watch mode
bun test:watch

# Run with coverage
bun test --coverage
```

### Test Coverage
The test suite covers:
- ✅ Basic value formatting
- ✅ Complex nested structures
- ✅ List and array handling
- ✅ Custom formatter system
- ✅ Built-in type handling (Date, URL, etc.)
- ✅ Multiline content processing
- ✅ Edge cases and error handling
- ✅ Formatter composition and precedence

## 🌐 Runtime Compatibility

LLML TypeScript supports:
- ✅ Node.js 16+
- ✅ Bun runtime
- ✅ Deno (with npm imports)
- ✅ Modern browsers (ES2020+)

## 🏗️ Development

```bash
# Clone the repo
git clone https://github.com/yourusername/llml.git
cd llml/ts

# Install dependencies
bun install

# Run tests
bun test

# Run tests in watch mode
bun test:watch

# Type checking
bun run tsc --noEmit
```

## 📖 API Reference

### `llml(data, formatters?)`

**Parameters:**
- `data: unknown` - The data to convert to markup
- `formatters?: Formatters` - Optional custom formatters (defaults to `vibeXML()`)

**Types:**
```typescript
type Predicate = (value: unknown) => boolean;
type Formatter = (
  value: unknown,
  llml: (data: unknown, formatters: Formatters) => string,
  formatters: Formatters
) => string;
type Formatters = Iterable<[Predicate, Formatter]>;
```

**Returns:** `string` - The formatted markup

### `vibeXML(options?)`

Creates the default VibeXML formatters with optional customization:

**Parameters:**
- `options?: { formatters?: Formatters }` - Custom formatters to merge with defaults

**Examples:**
```typescript
// Simple usage with defaults
llml({ name: "John" })
// → <name>John</name>

// Custom formatters
class User {
  constructor(public name: string) {}
}

const customFormatters = new Map();
customFormatters.set(
  (v: unknown): v is User => v instanceof User,
  (v: User) => `User: ${v.name}`
);

const formatters = vibeXML({ formatters: customFormatters });
llml({ admin: new User("Alice") }, formatters)
// → <admin>User: Alice</admin>

// Complex data
llml({
  users: [
    { name: "Alice", age: 30 },
    { name: "Bob", age: 25 }
  ]
})
// → <users>
//     <users-1>
//       <name>Alice</name>
//       <age>30</age>
//     </users-1>
//     <users-2>
//       <name>Bob</name>
//       <age>25</age>
//     </users-2>
//   </users>
```

For detailed information on creating custom formatters, see our [Formatters Guide](docs/formatters.md).

## 🤝 Contributing

We love contributions! Whether it's:
- 🐛 Bug reports
- 💡 Feature requests
- 📝 Documentation improvements
- 🔧 Code contributions

Check out our [contribution guidelines](../CONTRIBUTING.md) to get started.

## 📄 License

MIT License - see [LICENSE](../LICENSE) file for details.

## 🌟 Star History

If LLML makes your life easier, give us a star! ⭐

---

**Made with ❤️ for the TypeScript community**

*LLML: Because beautiful markup shouldn't be hard.*
