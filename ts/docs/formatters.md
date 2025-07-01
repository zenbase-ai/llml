# LLML Formatters Guide (TypeScript)

Learn how to use and extend LLML's powerful formatter system to customize data serialization in TypeScript.

## Overview

LLML uses a **formatter system** to convert different data types into XML-like markup. Each formatter consists of:

1. **Predicate function**: Tests if a value matches a specific type
2. **Format function**: Converts the value to a string representation

When you call `llml(data)`, it iterates through the formatters in order until it finds one whose predicate returns `true`, then uses that formatter's format function.

```typescript
import { llml } from "@zenbase/llml";

// Uses default SwagXML formatters
const result = llml({
    name: "Alice",
    age: 30,
    active: true,
    tasks: ["code", "test", "deploy"]
});
```

## How Formatters Work

### The Formatter Structure

Each formatter is a tuple of `[predicate, formatFunction]`:

```typescript
import type { FormatFunction, Predicate } from "../src";

const isString = (v: unknown): v is string =>
    typeof v === "string" && !v.includes("\n");

const formatString: FormatFunction = (value: string) =>
    value.trim();
```

### Default SwagXML Formatters

The `swagXML()` function creates a formatter Map with built-in formatters:

```typescript
import { llml, swagXML } from "../src";

// Get the default formatter map
const formatters = swagXML;
console.log(formatters); // Map(8) { [Function: isString] => [Function: formatString], ... }

// Use custom formatters
const result = llml(data, formatters);
```

The default formatters handle:
- **Strings**: Simple text (multiline strings handled specially)
- **Numbers**: `number` type
- **Booleans**: `true`/`false`
- **Undefined**: Converts to empty string `""`
- **Null**: Converts to `"null"`
- **Dates**: `Date` objects
- **Arrays**: `Array` type with numbered items
- **Objects**: Plain objects and classes

## Extending Formatters

### Method 1: Modify Existing Formatter Map

Replace or add formatters to the default map:

```typescript
import { llml, swagXML } from "@zenbase/llml";
import { isDate } from "@zenbase/llml/formatters/swag-xml";

const formatDateCustom = (value: Date) =>
    value.toISOString().split('T')[0]; // YYYY-MM-DD format

// Get default formatters and override date formatting
const formatters = swagXML();
formatters.set(isDate, formatDateCustom);

// Use with custom date formatting
const data = { created: new Date("2023-12-25T10:30:00Z") };
const result = llml(data, formatters);
// Output: <created>2023-12-25</created>
```

### Method 2: Create New Formatter Map

Put your custom formatters first (higher priority):

```typescript
import { llml, swagXML } from "@zenbase/llml";

class User {
    constructor(public name: string, public email: string) {}
}

const isUser = (v: unknown): v is User => v instanceof User;

const formatUser = (value: User) => `${value.name} (${value.email})`;

// Create new map with custom formatter first
const formatters = swagXML({
    formatters: [isUser, formatUser]
})

const data = { admin: new User("Alice", "alice@example.com") };
const result = llml(data, formatters);
// Output: <admin>Alice (alice@example.com)</admin>
```

### Method 3: Using swagXML Options

Use the built-in custom formatter support:

```typescript
const formatters = swagXML({ formatters: [[isUser, formatUser]] });
const result = llml(data, formatters);
```

## Custom Formatter Examples

### URL and Email Formatting

```typescript
const isEmail = (v: unknown): v is string => {
    if (typeof v !== "string") return false;
    return /^[^@]+@[^@]+\.[^@]+$/.test(v);
};

const formatEmail = (value: string) => `mailto:${value}`;

const isURL = (v: unknown): v is string => {
    if (typeof v !== "string") return false;
    return v.startsWith('http://') || v.startsWith('https://');
};

const formatURL = (value: string) => `[${value}]`;

// These take precedence over the builtin string formatter
const formatters = swagXML({
    formatters: [[isEmail, formatEmail], [isURL, formatURL]]
});

const data = {
    contact: "alice@example.com",
    website: "https://example.com",
    name: "Alice"  // Regular string, uses default formatter
};

const result = llml(data, formatters);
// Output:
// <contact>mailto:alice@example.com</contact>
// <website>[https://example.com]</website>
// <name>Alice</name>
```

### Formatting with .toString()

```typescript
class Temperature {
    constructor(public celsius: number) {}

    get fahrenheit(): number {
        return (this.celsius * 9/5) + 32;
    }

    toString() {
        return `${this.celsius}°C (${this.fahrenheit.toFixed(1)}°F)`;
    }
}

class Distance {
    constructor(public meters: number) {}

    get kilometers(): number {
        return this.meters / 1000;
    }
}

const data = {
    temperature: new Temperature(25),
    distance: new Distance(1500)
};

const result = llml(data);
// Output:
// <temperature>25°C (77.0°F)</temperature>
// <distance>1.50 km</distance>
```

## Advanced Usage

### Recursive Formatting

Format functions receive the `llml` function for recursive processing:

```typescript
import type { FormatterMap } from "../src";

class Project {
    constructor(public name: string, public tasks: string[]) {}
}

const isProject = (v: unknown): v is Project => v instanceof Project;

const formatProject = (
    value: Project,
    llml: (data: unknown, formatters: FormatterMap) => string,
    formatters: FormatterMap
) => {
    // Use the llml function recursively for the tasks
    const tasksFormatted = llml({ tasks: value.tasks }, formatters);
    return `Project: ${value.name}\n${tasksFormatted}`;
};

const formatters = new Map();
formatters.set(isProject, formatProject);

// Add default formatters
const defaultFormatters = swagXML();
for (const [predicate, formatter] of defaultFormatters) {
    formatters.set(predicate, formatter);
}

const project = new Project("LLML", ["implement", "test", "document"]);
const result = llml({ current: project }, formatters);
// Output:
// <current>Project: LLML
// <tasks>
//   <tasks-1>implement</tasks-1>
//   <tasks-2>test</tasks-2>
//   <tasks-3>document</tasks-3>
// </tasks></current>
```

### Conditional Formatting

Create formatters that behave differently based on context:

```typescript
const isSensitiveString = (v: unknown): v is string => {
    if (typeof v !== "string") return false;
    const sensitiveWords = ["password", "token", "secret", "key"];
    return sensitiveWords.some(word => v.toLowerCase().includes(word));
};

const formatSensitive = (value: string) => "[REDACTED]";

const isLargeNumber = (v: unknown): v is number =>
    typeof v === "number" && Math.abs(v) >= 1000000;

const formatLargeNumber = (value: number) => {
    if (value >= 1000000) {
        return `${(value/1000000).toFixed(1)}M`;
    } else if (value >= 1000) {
        return `${(value/1000).toFixed(1)}K`;
    }
    return String(value);
};

const formatters = new Map();
formatters.set(isSensitiveString, formatSensitive);
formatters.set(isLargeNumber, formatLargeNumber);

// Add default formatters
const defaultFormatters = swagXML();
for (const [predicate, formatter] of defaultFormatters) {
    formatters.set(predicate, formatter);
}

const data = {
    apiKey: "secret_12345",
    userCount: 1500000,
    name: "Alice"
};

const result = llml(data, formatters);
// Output:
// <apiKey>[REDACTED]</apiKey>
// <userCount>1.5M</userCount>
// <name>Alice</name>
```

## Order Matters

Formatters are processed in order - the first matching predicate wins:

```typescript
const isPositiveNumber = (v: unknown): v is number =>
    typeof v === "number" && v > 0;

const formatPositive = (value: number) => `+${value}`;

const isNumber = (v: unknown): v is number => typeof v === "number";

const formatNumber = (value: number) => String(value);

// Correct order: more specific first
const formatters = new Map();
formatters.set(isPositiveNumber, formatPositive);  // More specific
formatters.set(isNumber, formatNumber);             // Less specific

const result = llml({ score: 85 }, formatters);
// Output: <score>+85</score>

// Wrong order: general formatter catches everything
const wrongFormatters = new Map();
wrongFormatters.set(isNumber, formatNumber);             // Too general, catches everything
wrongFormatters.set(isPositiveNumber, formatPositive);  // Never reached!

const wrongResult = llml({ score: 85 }, wrongFormatters);
// Output: <score>85</score> (positive formatter never used)
```

## Best Practices

### 1. Specific Before General
Always put more specific formatters before general ones.

### 2. Use Type Guards
Write clear, efficient predicate functions:

```typescript
const isNonEmptyString = (v: unknown): v is string =>
    typeof v === "string" && v.length > 0;

const isEmailAddress = (v: unknown): v is string => {
    if (typeof v !== "string") return false;
    return v.includes("@") && v.split("@")[1]?.includes(".");
};
```

### 3. Handle Edge Cases
Make your formatters robust:

```typescript
const formatSafeURL = (value: string) => {
    if (!value) {
        return "[empty URL]";
    }

    try {
        // Validate URL format
        let url = value;
        if (!url.startsWith('http://') && !url.startsWith('https://')) {
            url = `https://${url}`;
        }
        return `[${url}]`;
    } catch (error) {
        return `[invalid URL: ${value}]`;
    }
};
```

### 4. Document Your Formatters
Make it clear what each formatter does:

```typescript
interface CurrencyAmount {
    amount: number;
    currency: 'USD' | 'EUR' | 'GBP';
}

/** Detects numeric values that should be formatted as currency. */
const isCurrencyAmount = (v: unknown): v is CurrencyAmount =>
    typeof v === "object" && v !== null &&
    "amount" in v && "currency" in v;

/** Formats currency amounts with appropriate symbol and precision. */
const formatCurrency = (value: CurrencyAmount) => {
    const symbols = { USD: '$', EUR: '€', GBP: '£' };
    const symbol = symbols[value.currency] || '$';
    return `${symbol}${value.amount.toFixed(2)}`;
};
```

## Complete Example: Custom Data Types

Here's a complete example showing how to handle custom data types:

```typescript
import { llml, swagXML } from "../src";
import type { FormatterMap } from "../src";

class BlogPost {
    constructor(
        public title: string,
        public author: Author,
        public published: Date,
        public tags: string[]
    ) {}
}

class Author {
    constructor(public name: string, public email: string) {}
}

// Custom formatters
const isBlogPost = (v: unknown): v is BlogPost => v instanceof BlogPost;

const formatBlogPost = (
    value: BlogPost,
    llml: (data: unknown, formatters: FormatterMap) => string,
    formatters: FormatterMap
) => {
    // Format the blog post with nested data
    const postData = {
        title: value.title,
        author: value.author,
        published: value.published,
        tags: value.tags
    };
    return llml(postData, formatters);
};

const isAuthor = (v: unknown): v is Author => v instanceof Author;

const formatAuthor = (value: Author) => `${value.name} <${value.email}>`;

const isDate = (v: unknown): v is Date => v instanceof Date;

const formatBlogDateTime = (value: Date) =>
    value.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
    });

// Create formatter map
const formatters = new Map();
formatters.set(isBlogPost, formatBlogPost);
formatters.set(isAuthor, formatAuthor);
formatters.set(isDate, formatBlogDateTime);

// Add default formatters
const defaultFormatters = swagXML();
for (const [predicate, formatter] of defaultFormatters) {
    formatters.set(predicate, formatter);
}

// Use the formatters
const post = new BlogPost(
    "LLML Formatters Guide",
    new Author("Alice", "alice@example.com"),
    new Date("2023-12-25"),
    ["llml", "typescript", "formatters"]
);

const result = llml({ post }, formatters);
console.log(result);

// Output:
// <post>
//   <title>LLML Formatters Guide</title>
//   <author>Alice <alice@example.com></author>
//   <published>December 25, 2023</published>
//   <tags>
//     <tags-1>llml</tags-1>
//     <tags-2>typescript</tags-2>
//     <tags-3>formatters</tags-3>
//   </tags>
// </post>
```

## Helper Functions

For easier formatter management, you can create utility functions:

```typescript
import type { FormatterMap, Predicate, FormatFunction } from "../src";
import { swagXML } from "../src";

/** Creates a new formatter map with custom formatters taking precedence */
export function createFormatters(
    customFormatters: Map<Predicate, FormatFunction>
): FormatterMap {
    const formatters = new Map(customFormatters);

    // Add default formatters that don't conflict
    const defaultFormatters = swagXML();
    for (const [predicate, formatter] of defaultFormatters) {
        if (!formatters.has(predicate)) {
            formatters.set(predicate, formatter);
        }
    }

    return formatters;
}

/** Adds a formatter to an existing map */
export function addFormatter<T>(
    formatters: FormatterMap,
    predicate: (v: unknown) => v is T,
    formatter: FormatFunction
): FormatterMap {
    const newFormatters = new Map(formatters);
    newFormatters.set(predicate, formatter);
    return newFormatters;
}

// Usage:
const customFormatters = new Map();
customFormatters.set(isUser, formatUser);
customFormatters.set(isTemperature, formatTemperature);

const formatters = createFormatters(customFormatters);
```

## Next Steps

- Explore the built-in formatters in `src/renderers/swag/`
- Create your own formatter library for common data types
- Consider performance implications for large datasets
- Share useful formatters with the community!

The formatter system makes LLML incredibly flexible while maintaining the clean, readable output format. Happy formatting! 🎉
