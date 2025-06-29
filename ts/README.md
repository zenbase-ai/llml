# 🚀 LLML TypeScript - Lightweight Language Markup Language

**The most elegant way to generate structured text in TypeScript/JavaScript.**

LLML transforms your data into beautifully formatted XML-like markup with zero fuss and maximum flexibility. Perfect for prompt engineering, configuration generation, and structured document creation.

## ⚡ Quick Start

```typescript
import { llml } from './src';

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
- **🔧 Zero Dependencies**: Only requires `dedent` for multiline handling
- **⚙️ Strict Mode**: Control nested property prefixes with `strict` option

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

### Automatic Key Conversion
LLML automatically converts all keys to kebab-case for consistency:

```typescript
// camelCase, snake_case, and spaces all become kebab-case
console.log(llml({ 
  userName: "Alice",           // camelCase
  user_age: 30,               // snake_case
  "user email": "alice@example.com"  // spaces
}));
// Output:
// <user-name>Alice</user-name>
// <user-age>30</user-age>
// <user-email>alice@example.com</user-email>

// Handles complex cases like acronyms
console.log(llml({ XMLHttpRequest: "api", HTMLElement: "dom" }));
// Output:
// <xml-http-request>api</xml-http-request>
// <html-element>dom</html-element>
```

### Prefix Support
```typescript
// Add prefix to all keys
console.log(llml({ message: "Hello" }, { prefix: "app" }));
// Output: <app-message>Hello</app-message>
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
//   <user-level>beginner</user-level>
//   <topic>TypeScript programming</topic>
// </context>

// Example with strict mode
console.log(llml({config: {debug: true, timeout: 30}}, {strict: true}));
// Output: <config>
//           <config-debug>true</config-debug>
//           <config-timeout>30</config-timeout>
//         </config>

// Example with strict mode disabled (default)
console.log(llml({config: {debug: true, timeout: 30}}, {strict: false}));
// Output: <config>
//           <debug>true</debug>
//           <timeout>30</timeout>
//         </config>
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
- ✅ Kebab-case key conversion (camelCase, snake_case, spaces)
- ✅ Multiline content processing
- ✅ Prefix functionality
- ✅ Indentation control
- ✅ Edge cases and error handling

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

### `llml(data, options?)`

**Parameters:**
- `data: any` - The data to convert to markup
- `options?: LLMLOptions` - Optional configuration

**Options:**
```typescript
interface LLMLOptions {
  indent?: string;  // Custom indentation string (default: "")
  prefix?: string;  // Prefix for all keys (default: "")
  strict?: boolean; // Include parent key prefixes in nested objects (default: false)
}
```

**Returns:** `string` - The formatted markup

**Examples:**
```typescript
// Simple usage
llml({ name: "John" })
// → <name>John</name>

// With options
llml({ name: "John" }, { prefix: "user", indent: "  " })
// → <user-name>John</user-name>

// Complex data
llml({
  users: [
    { name: "Alice", age: 30 },
    { name: "Bob", age: 25 }
  ]
})
// → <users-list>
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
