# LLML - Lightweight Language Markup Language

Convert data structures into human-readable, XML-like markup across multiple languages.

LLML is a data serialization library available in **Python**, **TypeScript/JavaScript**, **Rust**, and **Go** that transforms nested data structures (dictionaries/objects, lists/arrays, primitives) into well-formatted, XML-like markup.

📋 [Full Technical Specification](.cursor/rules/spec.mdc)

```python
# Python
from zenbase_llml import llml

llml({name: "Alice", age: 30})
# Output: <name>Alice</name>
#         <age>30</age>

llml({instructions: "You are a helpful assistant.", rules: ["be concise", "be helpful", "be accurate"]})
# Output: <instructions>You are a helpful assistant.</instructions>
#         <rules-list>
#           <rules-1>be concise</rules-1>
#           <rules-2>be helpful</rules-2>
#           <rules-3>be accurate</rules-3>
#         </rules-list>
```

<details>
<summary><b>Quick Examples in TS/JS/Rust/Go</b></summary>
<br>

```typescript
// TypeScript/JavaScript
import { llml } from '@zenbase/llml';

const agentPrompt = llml({
    role: "DevOps automation agent",
    context: {
        environment: "production",
        awsRegion: "us-east-1",
        services: ["web-api", "worker-queue", "database"],
        lastDeployment: "2024-01-15T10:30:00Z"
    },
    instructions: "Execute deployment workflow with safety checks",
    workflows: {
        deploy: [
            "Run pre-deployment health checks",
            "Create backup of current state",
            "Deploy to canary instance (5% traffic)",
            "Monitor metrics for 10 minutes",
            "If healthy, proceed to full deployment",
            "If issues detected, automatic rollback"
        ],
        rollback: [
            "Stop new traffic to affected services",
            "Restore from latest backup",
            "Verify service health",
            "Send notification to ops channel"
        ]
    },
    safetyRules: [
        "Never skip health checks",
        "Always maintain 99.9% uptime SLA",
        "Require manual approval for database changes"
    ]
});
// Output:
// <role>DevOps automation agent</role>
// <context>
//   <context-environment>production</context-environment>
//   <context-aws-region>us-east-1</context-aws-region>
//   <context-services-list>
//     <context-services-1>web-api</context-services-1>
//     <context-services-2>worker-queue</context-services-2>
//     <context-services-3>database</context-services-3>
//   </context-services-list>
//   <context-last-deployment>2024-01-15T10:30:00Z</context-last-deployment>
// </context>
// <instructions>Execute deployment workflow with safety checks</instructions>
// <workflows>
//   <workflows-deploy-list>
//     <workflows-deploy-1>Run pre-deployment health checks</workflows-deploy-1>
//     <workflows-deploy-2>Create backup of current state</workflows-deploy-2>
//     <workflows-deploy-3>Deploy to canary instance (5% traffic)</workflows-deploy-3>
//     <workflows-deploy-4>Monitor metrics for 10 minutes</workflows-deploy-4>
//     <workflows-deploy-5>If healthy, proceed to full deployment</workflows-deploy-5>
//     <workflows-deploy-6>If issues detected, automatic rollback</workflows-deploy-6>
//   </workflows-deploy-list>
//   <workflows-rollback-list>
//     <workflows-rollback-1>Stop new traffic to affected services</workflows-rollback-1>
//     <workflows-rollback-2>Restore from latest backup</workflows-rollback-2>
//     <workflows-rollback-3>Verify service health</workflows-rollback-3>
//     <workflows-rollback-4>Send notification to ops channel</workflows-rollback-4>
//   </workflows-rollback-list>
// </workflows>
// <safety-rules-list>
//   <safety-rules-1>Never skip health checks</safety-rules-1>
//   <safety-rules-2>Always maintain 99.9% uptime SLA</safety-rules-2>
//   <safety-rules-3>Require manual approval for database changes</safety-rules-3>
// </safety-rules-list>
```

```rust
// Rust
use zenbase_llml::llml;

let rag_prompt = llml(&json!({
    "system": "You are a helpful documentation assistant",
    "instructions": "Answer questions based on the provided documentation context",
    "documents": [
        {
            "title": "API Authentication Guide",
            "content": "Our API uses OAuth 2.0 for authentication...",
            "relevance_score": 0.95
        },
        {
            "title": "Rate Limiting Documentation",
            "content": "API calls are limited to 1000 requests per hour...",
            "relevance_score": 0.82
        }
    ],
    "user_query": "How do I authenticate with your API?",
    "constraints": [
        "Only use information from the provided documents",
        "Cite the document title when referencing information",
        "If information is not available, explicitly state so"
    ]
}));
// Output:
// <system>You are a helpful documentation assistant</system>
// <instructions>Answer questions based on the provided documentation context</instructions>
// <documents-list>
//   <documents-1>
//     <documents-1-title>API Authentication Guide</documents-1-title>
//     <documents-1-content>Our API uses OAuth 2.0 for authentication...</documents-1-content>
//     <documents-1-relevance-score>0.95</documents-1-relevance-score>
//   </documents-1>
//   <documents-2>
//     <documents-2-title>Rate Limiting Documentation</documents-2-title>
//     <documents-2-content>API calls are limited to 1000 requests per hour...</documents-2-content>
//     <documents-2-relevance-score>0.82</documents-2-relevance-score>
//   </documents-2>
// </documents-list>
// <user-query>How do I authenticate with your API?</user-query>
// <constraints-list>
//   <constraints-1>Only use information from the provided documents</constraints-1>
//   <constraints-2>Cite the document title when referencing information</constraints-2>
//   <constraints-3>If information is not available, explicitly state so</constraints-3>
// </constraints-list>
```

```go
// Go
package main

import (
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

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
// Output:
// <system>You are a helpful documentation assistant</system>
// <instructions>Answer questions based on the provided documentation context</instructions>
// <documents-list>
//   <documents-1>
//     <documents-1-title>API Authentication Guide</documents-1-title>
//     <documents-1-content>Our API uses OAuth 2.0 for authentication...</documents-1-content>
//     <documents-1-relevance-score>0.95</documents-1-relevance-score>
//   </documents-1>
//   <documents-2>
//     <documents-2-title>Rate Limiting Documentation</documents-2-title>
//     <documents-2-content>API calls are limited to 1000 requests per hour...</documents-2-content>
//     <documents-2-relevance-score>0.82</documents-2-relevance-score>
//   </documents-2>
// </documents-list>
// <user-query>How do I authenticate with your API?</user-query>
// <constraints-list>
//   <constraints-1>Only use information from the provided documents</constraints-1>
//   <constraints-2>Cite the document title when referencing information</constraints-2>
//   <constraints-3>If information is not available, explicitly state so</constraints-3>
// </constraints-list>
```
</details>

## Benefits

1. 🏗️ **Structured Prompt Engineering**: Replace fragile string concatenation and template management with systematic data-to-markup conversion.
2. 📊 **Complex Data Simplification**: Transform nested configuration objects, user data, and multi-part instructions into clean, readable markup without manual formatting.
3. 🔀 **Cross-Language Reliability**: Identical output across Python, TypeScript, Rust, and Go means your prompt generation works the same regardless of your backend language.
4. 🛠️ **Maintainable Prompt Templates**: Stop maintaining brittle template strings that break when data structures change. With LLML, your prompts are just data structures—add new fields, reorganize objects, or change array contents, and the markup automatically updates with proper formatting and numbering.
5. 🤖 **LLM-Optimized Output**: LLML's XML-like format with clear tag boundaries and numbered list items (`<rules-1>`, `<rules-2>`) reduces prompt ambiguity and reduces attentional load. Language models can reliably identify and reference specific sections, improving response accuracy and reducing hallucinations in complex, multi-part prompts.

## Table of Contents

- [LLML - Lightweight Language Markup Language](#llml---lightweight-language-markup-language)
  - [Benefits](#benefits)
  - [Table of Contents](#table-of-contents)
  - [How It Works](#how-it-works)
  - [Features](#features)
  - [Quick Start](#quick-start)
    - [Python](#python)
    - [TypeScript/JavaScript](#typescriptjavascript)
    - [Rust](#rust)
    - [Go](#go)
  - [Prerequisites](#prerequisites)
  - [Installation \& Setup](#installation--setup)
    - [Python Project](#python-project)
    - [TypeScript Project](#typescript-project)
    - [Rust Project](#rust-project)
    - [Go Project](#go-project)
  - [Linting](#linting)
  - [Running Tests](#running-tests)
  - [Development](#development)
  - [API Reference](#api-reference)
  - [License](#license)

## How It Works

LLML transforms data using these core rules:

1. **Simple Values**: `{key: "value"}` → `<key>value</key>`
2. **Lists/Arrays**: `{items: ["a", "b"]}` → `<items-list><items-1>a</items-1><items-2>b</items-2></items-list>`
3. **Nested Objects**: `{config: {debug: true}}` → `<config><config-debug>true</config-debug></config>`
4. **Key Normalization**: All keys are converted to kebab-case automatically
5. **Empty Values**: Empty objects `{}` and arrays `[]` return empty strings


## Features

- 🔄 **Automatic kebab-case conversion**: Transform `user_name` → `user-name`, `maxRetries` → `max-retries`
- 📝 **Smart list formatting**: Arrays become `<items-list><items-1>first</items-1><items-2>second</items-2></items-list>`
- 🔁 **Recursive nested structures**: Objects within objects maintain proper hierarchy with compound keys
- 📐 **Configurable indentation**: Custom spacing for nested elements
- 📄 **Multiline content support**: Preserves line breaks with proper indentation
- 🏷️ **Optional prefix namespacing**: Add prefixes to all generated tags
- ⚡ **Zero configuration**: Works out of the box with sensible defaults
- 🔀 **Consistent cross-language output**: Identical results across Python, TypeScript, Rust, and Go

## Quick Start

### Python

```bash
uv install zenbase-llml # or pip, rye, poetry, etc.
```

```python
from zenbase_llml import llml

# Simple example
llml(task="analyze", content="customer feedback")
# Output: <task>analyze</task>
#         <content>customer feedback</content>

# List handling
llml(rules=["be concise", "be helpful", "be accurate"])
# Output: <rules-list>
#           <rules-1>be concise</rules-1>
#           <rules-2>be helpful</rules-2>
#           <rules-3>be accurate</rules-3>
#         </rules-list>

# Example 1: Data Extraction
extraction_prompt = llml(
    task="Extract key information from customer feedback",
    instructions="Identify and categorize customer sentiments and specific issues mentioned",
    rules=[
        "Classify sentiment as positive, negative, or neutral",
        "Extract specific product features mentioned",
        "Identify any requested improvements or fixes",
        "Note any comparisons to competitors"
    ],
    output_format={
        "sentiment": "positive/negative/neutral",
        "features_mentioned": ["list of features"],
        "issues": ["list of problems"],
        "improvements": ["list of suggestions"]
    }
)
# Output:
# <task>Extract key information from customer feedback</task>
# <instructions>Identify and categorize customer sentiments and specific issues mentioned</instructions>
# <rules-list>
#   <rules-1>Classify sentiment as positive, negative, or neutral</rules-1>
#   <rules-2>Extract specific product features mentioned</rules-2>
#   <rules-3>Identify any requested improvements or fixes</rules-3>
#   <rules-4>Note any comparisons to competitors</rules-4>
# </rules-list>
# <output-format>
#   <output-format-sentiment>positive/negative/neutral</output-format-sentiment>
#   <output-format-features-mentioned-list>
#     <output-format-features-mentioned-1>list of features</output-format-features-mentioned-1>
#   </output-format-features-mentioned-list>
#   <output-format-issues-list>
#     <output-format-issues-1>list of problems</output-format-issues-1>
#   </output-format-issues-list>
#   <output-format-improvements-list>
#     <output-format-improvements-1>list of suggestions</output-format-improvements-1>
#   </output-format-improvements-list>
# </output-format>

# Example 2: RAG Chatbot
rag_prompt = llml(
    system="You are a helpful documentation assistant",
    instructions="Answer questions based on the provided documentation context",
    documents=[
        {
            "title": "API Authentication Guide",
            "content": "Our API uses OAuth 2.0 for authentication...",
            "relevance_score": 0.95
        },
        {
            "title": "Rate Limiting Documentation",
            "content": "API calls are limited to 1000 requests per hour...",
            "relevance_score": 0.82
        }
    ],
    user_query="How do I authenticate with your API?",
    constraints=[
        "Only use information from the provided documents",
        "Cite the document title when referencing information",
        "If information is not available, explicitly state so"
    ]
)
# Output:
# <system>You are a helpful documentation assistant</system>
# <instructions>Answer questions based on the provided documentation context</instructions>
# <documents-list>
#   <documents-1>
#     <documents-1-title>API Authentication Guide</documents-1-title>
#     <documents-1-content>Our API uses OAuth 2.0 for authentication...</documents-1-content>
#     <documents-1-relevance-score>0.95</documents-1-relevance-score>
#   </documents-1>
#   <documents-2>
#     <documents-2-title>Rate Limiting Documentation</documents-2-title>
#     <documents-2-content>API calls are limited to 1000 requests per hour...</documents-2-content>
#     <documents-2-relevance-score>0.82</documents-2-relevance-score>
#   </documents-2>
# </documents-list>
# <user-query>How do I authenticate with your API?</user-query>
# <constraints-list>
#   <constraints-1>Only use information from the provided documents</constraints-1>
#   <constraints-2>Cite the document title when referencing information</constraints-2>
#   <constraints-3>If information is not available, explicitly state so</constraints-3>
# </constraints-list>

# Example 3: AI Agent with Workflows
agent_prompt = llml(
    role="DevOps automation agent",
    context={
        "environment": "production",
        "aws_region": "us-east-1",
        "services": ["web-api", "worker-queue", "database"],
        "last_deployment": "2024-01-15T10:30:00Z"
    },
    instructions="Execute deployment workflow with safety checks",
    workflows={
        "deploy": [
            "Run pre-deployment health checks",
            "Create backup of current state",
            "Deploy to canary instance (5% traffic)",
            "Monitor metrics for 10 minutes",
            "If healthy, proceed to full deployment",
            "If issues detected, automatic rollback"
        ],
        "rollback": [
            "Stop new traffic to affected services",
            "Restore from latest backup",
            "Verify service health",
            "Send notification to ops channel"
        ]
    },
    safety_rules=[
        "Never skip health checks",
        "Always maintain 99.9% uptime SLA",
        "Require manual approval for database changes"
    ]
)
print(agent_prompt)
# Output:
# <role>DevOps automation agent</role>
# <context>
#   <context-environment>production</context-environment>
#   <context-aws-region>us-east-1</context-aws-region>
#   <context-services-list>
#     <context-services-1>web-api</context-services-1>
#     <context-services-2>worker-queue</context-services-2>
#     <context-services-3>database</context-services-3>
#   </context-services-list>
#   <context-last-deployment>2024-01-15T10:30:00Z</context-last-deployment>
# </context>
# <instructions>Execute deployment workflow with safety checks</instructions>
# <workflows>
#   <workflows-deploy-list>
#     <workflows-deploy-1>Run pre-deployment health checks</workflows-deploy-1>
#     <workflows-deploy-2>Create backup of current state</workflows-deploy-2>
#     <workflows-deploy-3>Deploy to canary instance (5% traffic)</workflows-deploy-3>
#     <workflows-deploy-4>Monitor metrics for 10 minutes</workflows-deploy-4>
#     <workflows-deploy-5>If healthy, proceed to full deployment</workflows-deploy-5>
#     <workflows-deploy-6>If issues detected, automatic rollback</workflows-deploy-6>
#   </workflows-deploy-list>
#   <workflows-rollback-list>
#     <workflows-rollback-1>Stop new traffic to affected services</workflows-rollback-1>
#     <workflows-rollback-2>Restore from latest backup</workflows-rollback-2>
#     <workflows-rollback-3>Verify service health</workflows-rollback-3>
#     <workflows-rollback-4>Send notification to ops channel</workflows-rollback-4>
#   </workflows-rollback-list>
# </workflows>
# <safety-rules-list>
#   <safety-rules-1>Never skip health checks</safety-rules-1>
#   <safety-rules-2>Always maintain 99.9% uptime SLA</safety-rules-2>
#   <safety-rules-3>Require manual approval for database changes</safety-rules-3>
# </safety-rules-list>
```

### TypeScript/JavaScript

```bash
pnpm add @zenbase/llml # or bun, npm, yarn, etc.
```

```typescript
import { llml } from '@zenbase/llml';

// Simple example
console.log(llml({ task: "analyze", content: "customer feedback" }));
// Output: <task>analyze</task>
//         <content>customer feedback</content>

// List handling
console.log(llml({ rules: ["be concise", "be helpful", "be accurate"] }));
// Output: <rules-list>
//           <rules-1>be concise</rules-1>
//           <rules-2>be helpful</rules-2>
//           <rules-3>be accurate</rules-3>
//         </rules-list>

// Example 1: Data Extraction
const extractionPrompt = llml({
    task: "Extract key information from customer feedback",
    instructions: "Identify and categorize customer sentiments and specific issues mentioned",
    rules: [
        "Classify sentiment as positive, negative, or neutral",
        "Extract specific product features mentioned",
        "Identify any requested improvements or fixes",
        "Note any comparisons to competitors"
    ],
    outputFormat: {
        sentiment: "positive/negative/neutral",
        featuresMentioned: ["list of features"],
        issues: ["list of problems"],
        improvements: ["list of suggestions"]
    }
});
// Output:
// <task>Extract key information from customer feedback</task>
// <instructions>Identify and categorize customer sentiments and specific issues mentioned</instructions>
// <rules-list>
//   <rules-1>Classify sentiment as positive, negative, or neutral</rules-1>
//   <rules-2>Extract specific product features mentioned</rules-2>
//   <rules-3>Identify any requested improvements or fixes</rules-3>
//   <rules-4>Note any comparisons to competitors</rules-4>
// </rules-list>
// <output-format>
//   <output-format-sentiment>positive/negative/neutral</output-format-sentiment>
//   <output-format-features-mentioned>list of features</output-format-features-mentioned>
//   <output-format-issues>list of problems</output-format-issues>
//   <output-format-improvements>list of suggestions</output-format-improvements>
// </output-format>

// Example 2: RAG Chatbot
const ragPrompt = llml({
    system: "You are a helpful documentation assistant",
    instructions: "Answer questions based on the provided documentation context",
    documents: [
        {
            title: "API Authentication Guide",
            content: "Our API uses OAuth 2.0 for authentication...",
            relevanceScore: 0.95
        },
        {
            title: "Rate Limiting Documentation",
            content: "API calls are limited to 1000 requests per hour...",
            relevanceScore: 0.82
        }
    ],
    userQuery: "How do I authenticate with your API?",
    constraints: [
        "Only use information from the provided documents",
        "Cite the document title when referencing information",
        "If information is not available, explicitly state so"
    ]
});
// Output:
// <system>You are a helpful documentation assistant</system>
// <instructions>Answer questions based on the provided documentation context</instructions>
// <documents-list>
//   <documents-1>
//     <documents-1-title>API Authentication Guide</documents-1-title>
//     <documents-1-content>Our API uses OAuth 2.0 for authentication...</documents-1-content>
//     <documents-1-relevance-score>0.95</documents-1-relevance-score>
//   </documents-1>
//   <documents-2>
//     <documents-2-title>Rate Limiting Documentation</documents-2-title>
//     <documents-2-content>API calls are limited to 1000 requests per hour...</documents-2-content>
//     <documents-2-relevance-score>0.82</documents-2-relevance-score>
//   </documents-2>
// </documents-list>
// <user-query>How do I authenticate with your API?</user-query>
// <constraints-list>
//   <constraints-1>Only use information from the provided documents</constraints-1>
//   <constraints-2>Cite the document title when referencing information</constraints-2>
//   <constraints-3>If information is not available, explicitly state so</constraints-3>
// </constraints-list>

// Example 3: AI Agent with Workflows
const agentPrompt = llml({
    role: "DevOps automation agent",
    context: {
        environment: "production",
        awsRegion: "us-east-1",
        services: ["web-api", "worker-queue", "database"],
        lastDeployment: "2024-01-15T10:30:00Z"
    },
    instructions: "Execute deployment workflow with safety checks",
    workflows: {
        deploy: [
            "Run pre-deployment health checks",
            "Create backup of current state",
            "Deploy to canary instance (5% traffic)",
            "Monitor metrics for 10 minutes",
            "If healthy, proceed to full deployment",
            "If issues detected, automatic rollback"
        ],
        rollback: [
            "Stop new traffic to affected services",
            "Restore from latest backup",
            "Verify service health",
            "Send notification to ops channel"
        ]
    },
    safetyRules: [
        "Never skip health checks",
        "Always maintain 99.9% uptime SLA",
        "Require manual approval for database changes"
    ]
});
// Output:
// <role>DevOps automation agent</role>
// <context>
//   <context-environment>production</context-environment>
//   <context-aws-region>us-east-1</context-aws-region>
//   <context-services-list>
//     <context-services-1>web-api</context-services-1>
//     <context-services-2>worker-queue</context-services-2>
//     <context-services-3>database</context-services-3>
//   </context-services-list>
//   <context-last-deployment>2024-01-15T10:30:00Z</context-last-deployment>
// </context>
// <instructions>Execute deployment workflow with safety checks</instructions>
// <workflows>
//   <workflows-deploy-list>
//     <workflows-deploy-1>Run pre-deployment health checks</workflows-deploy-1>
//     <workflows-deploy-2>Create backup of current state</workflows-deploy-2>
//     <workflows-deploy-3>Deploy to canary instance (5% traffic)</workflows-deploy-3>
//     <workflows-deploy-4>Monitor metrics for 10 minutes</workflows-deploy-4>
//     <workflows-deploy-5>If healthy, proceed to full deployment</workflows-deploy-5>
//     <workflows-deploy-6>If issues detected, automatic rollback</workflows-deploy-6>
//   </workflows-deploy-list>
//   <workflows-rollback-list>
//     <workflows-rollback-1>Stop new traffic to affected services</workflows-rollback-1>
//     <workflows-rollback-2>Restore from latest backup</workflows-rollback-2>
//     <workflows-rollback-3>Verify service health</workflows-rollback-3>
//     <workflows-rollback-4>Send notification to ops channel</workflows-rollback-4>
//   </workflows-rollback-list>
// </workflows>
// <safety-rules-list>
//   <safety-rules-1>Never skip health checks</safety-rules-1>
//   <safety-rules-2>Always maintain 99.9% uptime SLA</safety-rules-2>
//   <safety-rules-3>Require manual approval for database changes</safety-rules-3>
// </safety-rules-list>
```

### Rust

```bash
cargo add zenbase-llml # or cargo, cargo-binstall, etc.
```

```rust
use llml::llml;
use serde_json::json;

// Simple example
println!("{}", llml(&json!({ "task": "analyze", "content": "customer feedback" }), None));
// Output: <task>analyze</task>
//         <content>customer feedback</content>

// List handling
println!("{}", llml(&json!({ "rules": ["be concise", "be helpful", "be accurate"] }), None));
// Output: <rules-list>
//           <rules-1>be concise</rules-1>
//           <rules-2>be helpful</rules-2>
//           <rules-3>be accurate</rules-3>
//         </rules-list>

// Example 1: Data Extraction
let extraction_prompt = llml(&json!({
    "task": "Extract key information from customer feedback",
    "instructions": "Identify and categorize customer sentiments and specific issues mentioned",
    "rules": [
        "Classify sentiment as positive, negative, or neutral",
        "Extract specific product features mentioned",
        "Identify any requested improvements or fixes",
        "Note any comparisons to competitors"
    ],
    "output_format": {
        "sentiment": "positive/negative/neutral",
        "features_mentioned": ["list of features"],
        "issues": ["list of problems"],
        "improvements": ["list of suggestions"]
    }
}));
println!("{}", extraction_prompt);

// Example 2: RAG Chatbot
let rag_prompt = llml(&json!({
    "system": "You are a helpful documentation assistant",
    "instructions": "Answer questions based on the provided documentation context",
    "documents": [
        {
            "title": "API Authentication Guide",
            "content": "Our API uses OAuth 2.0 for authentication...",
            "relevance_score": 0.95
        },
        {
            "title": "Rate Limiting Documentation",
            "content": "API calls are limited to 1000 requests per hour...",
            "relevance_score": 0.82
        }
    ],
    "user_query": "How do I authenticate with your API?",
    "constraints": [
        "Only use information from the provided documents",
        "Cite the document title when referencing information",
        "If information is not available, explicitly state so"
    ]
}));
println!("{}", rag_prompt);

// Example 3: AI Agent with Workflows
let agent_prompt = llml(&json!({
    "role": "DevOps automation agent",
    "context": {
        "environment": "production",
        "aws_region": "us-east-1",
        "services": ["web-api", "worker-queue", "database"],
        "last_deployment": "2024-01-15T10:30:00Z"
    },
    "instructions": "Execute deployment workflow with safety checks",
    "workflows": {
        "deploy": [
            "Run pre-deployment health checks",
            "Create backup of current state",
            "Deploy to canary instance (5% traffic)",
            "Monitor metrics for 10 minutes",
            "If healthy, proceed to full deployment",
            "If issues detected, automatic rollback"
        ],
        "rollback": [
            "Stop new traffic to affected services",
            "Restore from latest backup",
            "Verify service health",
            "Send notification to ops channel"
        ]
    },
    "safety_rules": [
        "Never skip health checks",
        "Always maintain 99.9% uptime SLA",
        "Require manual approval for database changes"
    ]
}));
println!("{}", agent_prompt);
```

### Go

```bash
go install github.com/zenbase-ai/llml/go@latest
```

```go
package main

import (
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Simple example
llml.Sprintf(map[string]any{
  "task": "analyze",
  "content": "customer feedback",
})
// Output: <task>analyze</task>
//         <content>customer feedback</content>

// List handling
llml.Sprintf(map[string]any{
  "rules": []any{"be concise", "be helpful", "be accurate"},
})
// Output: <rules-list>
//           <rules-1>be concise</rules-1>
//           <rules-2>be helpful</rules-2>
//           <rules-3>be accurate</rules-3>
//         </rules-list>

// Example 1: Data Extraction
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

// Example 2: RAG Chatbot
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

  // Example 3: AI Agent with Workflows
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
```

## Prerequisites

This project uses [mise](https://mise.jdx.dev/) for tool management. First install the required development tools:

```bash
mise install
```

This will install the specific versions of:
- `just` (task runner)
- `python` (Python interpreter)
- `go` (Go compiler)
- `rust` (Rust compiler)
- `bun` (JavaScript runtime and package manager)

## Installation & Setup

After running `mise install`, set up the language-specific dependencies:

### Python Project
```bash
cd py/
uv install
```

### TypeScript Project
```bash
cd ts/
bun install
```

### Rust Project
```bash
cd rs/
cargo build
```

### Go Project
```bash
cd go/
go mod download
```

## Linting

```bash
just lint # run all linters
just lint ts # typescript linters
just lint py # python linters
just lint rs # rust linters
just lint go # go linters
```

## Running Tests

```bash
just test # run all tests
just test ts # typescript tests
just test py # python tests
just test rs # rust tests
just test go # go tests
```

## Development

This project uses:
- **Python**: `uv` for dependency management, `pytest` for testing, `ruff` for linting
- **TypeScript**: `bun` for runtime and package management, `vitest` for testing
- **Rust**: `cargo` for dependency management, `cargo test` for testing
- **Go**: `go` for dependency management, `go test` for testing
- **Just**: Task runner for cross-platform commands

## API Reference

All implementations provide the same core functionality:

- Convert dictionaries/objects to nested tag structures
- Transform lists/arrays into numbered item lists
- Handle primitives (strings, numbers, booleans) as tag content
- Automatic key formatting to kebab-case
- Configurable indentation and prefixing

See individual project READMEs for detailed API documentation:
- [Python API Documentation](py/README.md)
- [TypeScript API Documentation](ts/README.md)
- [Rust API Documentation](rs/README.md)
- [Go API Documentation](go/README.md)

## License

MIT
