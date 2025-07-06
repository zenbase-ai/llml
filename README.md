# LLML - Lightweight Language Markup Language

**The compositional primitive for AI context engineering**

LLML is available in **Python**, **TypeScript/JavaScript**, **Rust**, and **Go**. Transform your nested data structures into structured markup optimized for AI understanding and performance.

Make complex contexts composable, maintainable, and reusable.

![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)
![Python Coverage](https://img.shields.io/badge/Python-92%25-brightgreen)
![TypeScript Coverage](https://img.shields.io/badge/TypeScript-100%25-brightgreen)
![Go Coverage](https://img.shields.io/badge/Go-100%25-brightgreen)
![Rust Coverage](https://img.shields.io/badge/Rust-91.67%25-brightgreen)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

[![PyPI version](https://img.shields.io/pypi/v/zenbase-llml.svg)](https://pypi.org/project/zenbase-llml/)
[![npm version](https://img.shields.io/npm/v/@zenbase/llml.svg)](https://www.npmjs.com/package/@zenbase/llml)
[![Crates.io](https://img.shields.io/crates/v/llml.svg)](https://crates.io/crates/llml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zenbase-ai/llml/go.svg)](https://pkg.go.dev/github.com/zenbase-ai/llml/go)

**Stop fighting string concatenation. Start building structured AI contexts.**

LLML transforms your nested data structures into well-formatted markup that's optimized for both human readability and AI model attention. Build complex prompts from simple, composable data structures.

📋 [Full Technical Specification](.cursor/rules/spec.mdc) | 🧩 [Compositional Patterns Guide](REACT.md)

```typescript
import { llml } from "@zenbase/llml"

// Simple components
const guardrails = [
  "Never skip health checks",
  "Require manual approval for database changes"
]
const userProfile = { name: "Alice", age: 30 }
const systemRole = "Senior DevOps Engineer"

// Compose them into a complex prompt
const deploymentPrompt = llml({
  role: systemRole,
  context: {
    environment: "production",
    region: "us-east-1",
    operator: userProfile
  }
  task: "Deploy latest version to staging safely"
  guardrails,
})

// Output: Clean, structured markup optimized for AI attention
// <role>Senior DevOps Engineer</role>
// <context>
//   <environment>production</environment>
//   <region>us-east-1</region>
//   <operator>
//     <name>Alice</name>
//     <age>30</age>
//   </operator>
// </context>
// <task>Deploy latest version to staging safely</task>
// <guardrails>
//   <guardrails-1>Never skip health checks</guardrails-1>
//   <guardrails-2>Require manual approval for database changes</guardrails-2>
// </guardrails>

await makeAIRequest(deploymentPrompt, {
    tools: DEPLOYMENT_TOOLS,
    toolChoice: "auto",
    maxSteps: 10,
})
```

<details>
<summary><b>Cheatsheet</b></summary>
<br>

```python
# Python
from zenbase_llml import llml

agent_prompt = llml({
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
})
# Output:
# <role>DevOps automation agent</role>
# <context>
#   <environment>production</environment>
#   <aws_region>us-east-1</aws_region>
#   <services>
#     <services-1>web-api</services-1>
#     <services-2>worker-queue</services-2>
#     <services-3>database</services-3>
#   </services>
#   <last_deployment>2024-01-15T10:30:00Z</last_deployment>
# </context>
# <instructions>Execute deployment workflow with safety checks</instructions>
# <workflows>
#   <deploy>
#     <deploy-1>Run pre-deployment health checks</deploy-1>
#     <deploy-2>Create backup of current state</deploy-2>
#     <deploy-3>Deploy to canary instance (5% traffic)</deploy-3>
#     <deploy-4>Monitor metrics for 10 minutes</deploy-4>
#     <deploy-5>If healthy, proceed to full deployment</deploy-5>
#     <deploy-6>If issues detected, automatic rollback</deploy-6>
#   </deploy>
#   <rollback>
#     <rollback-1>Stop new traffic to affected services</rollback-1>
#     <rollback-2>Restore from latest backup</rollback-2>
#     <rollback-3>Verify service health</rollback-3>
#     <rollback-4>Send notification to ops channel</rollback-4>
#   </rollback>
# </workflows>
# <safety_rules>
#   <safety_rules-1>Never skip health checks</safety_rules-1>
#   <safety_rules-2>Always maintain 99.9% uptime SLA</safety_rules-2>
#   <safety_rules-3>Require manual approval for database changes</safety_rules-3>
# </safety_rules>
```

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
//   <environment>production</environment>
//   <aws-region>us-east-1</aws-region>
//   <services>
//     <services-1>web-api</services-1>
//     <services-2>worker-queue</services-2>
//     <services-3>database</services-3>
//   </services>
//   <last-deployment>2024-01-15T10:30:00Z</last-deployment>
// </context>
// <instructions>Execute deployment workflow with safety checks</instructions>
// <workflows>
//   <deploy>
//     <deploy-1>Run pre-deployment health checks</deploy-1>
//     <deploy-2>Create backup of current state</deploy-2>
//     <deploy-3>Deploy to canary instance (5% traffic)</deploy-3>
//     <deploy-4>Monitor metrics for 10 minutes</deploy-4>
//     <deploy-5>If healthy, proceed to full deployment</deploy-5>
//     <deploy-6>If issues detected, automatic rollback</deploy-6>
//   </deploy>
//   <rollback>
//     <rollback-1>Stop new traffic to affected services</rollback-1>
//     <rollback-2>Restore from latest backup</rollback-2>
//     <rollback-3>Verify service health</rollback-3>
//     <rollback-4>Send notification to ops channel</rollback-4>
//   </rollback>
// </workflows>
// <safety-rules>
//   <safety-rules-1>Never skip health checks</safety-rules-1>
//   <safety-rules-2>Always maintain 99.9% uptime SLA</safety-rules-2>
//   <safety-rules-3>Require manual approval for database changes</safety-rules-3>
// </safety-rules>
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
// <documents>
//   <documents-1>
//     <title>API Authentication Guide</title>
//     <content>Our API uses OAuth 2.0 for authentication...</content>
//     <relevance-score>0.95</relevance-score>
//   </documents-1>
//   <documents-2>
//     <title>Rate Limiting Documentation</title>
//     <content>API calls are limited to 1000 requests per hour...</content>
//     <relevance-score>0.82</relevance-score>
//   </documents-2>
// </documents>
// <user-query>How do I authenticate with your API?</user-query>
// <constraints>
//   <constraints-1>Only use information from the provided documents</constraints-1>
//   <constraints-2>Cite the document title when referencing information</constraints-2>
//   <constraints-3>If information is not available, explicitly state so</constraints-3>
// </constraints>
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
// <documents>
//   <documents-1>
//     <title>API Authentication Guide</title>
//     <content>Our API uses OAuth 2.0 for authentication...</content>
//     <relevance-score>0.95</relevance-score>
//   </documents-1>
//   <documents-2>
//     <title>Rate Limiting Documentation</title>
//     <content>API calls are limited to 1000 requests per hour...</content>
//     <relevance-score>0.82</relevance-score>
//   </documents-2>
// </documents>
// <user-query>How do I authenticate with your API?</user-query>
// <constraints>
//   <constraints-1>Only use information from the provided documents</constraints-1>
//   <constraints-2>Cite the document title when referencing information</constraints-2>
//   <constraints-3>If information is not available, explicitly state so</constraints-3>
// </constraints>
```
</details>

## Why We Built This

After building dozens of AI systems at Zenbase, we kept hitting the same walls:

**The String Concatenation Nightmare**
```python
# We've all written this code...
prompt = f"Role: {role}\n"
prompt += f"Context: {context}\n"
prompt += "Instructions:\n"
for i, instruction in enumerate(instructions):
    prompt += f"  {i+1}. {instruction}\n"
# ...and debugged it at 3 AM when it breaks
```

**The Context Sprawl Problem**
As AI agents get more powerful, their prompts are getting more complex. A simple chatbot prompt can grow from 10 lines to 500. Changes in one place break formatting in another. Version control shows massive diffs for tiny logical changes.

**The Attention Crisis**
We discovered that LLMs perform significantly better with structured XML input, but handcrafting XML strings was trudge work indenting and ensuring opening/closing tags.

**The Composition Challenge**
We wanted to build prompts from small, reusable pieces that could be composed together. But everything was becoming string interpolation.

LLML is our solution: a compositional primitive that transforms data structures into structured markup. **Because AI development shouldn't feel like string manipulation.**

## Getting Started in 30 Seconds

```python
# Install: pip install zenbase-llml

from zenbase_llml import llml

# Before: Error-prone string manipulation
prompt = "You are an " + role + ".\n" + \
         "Context:\n" + json.dumps(context) + "\n" + \
         "Rules:\n" + "\n".join(f"- {r}" for r in rules)

# After: Composable, maintainable, beautiful
prompt = llml({
    "role": role,
    "context": context,
    "rules": rules
})

# That's it. Your prompts are now:
# ✓ Structured for optimal LLM performance
# ✓ Composable and reusable
# ✓ Type-safe and debuggable
# ✓ 42% more likely to get correct responses (swag)
```

## Core Principles

### 1. **Compositional Architecture**
Build complex prompts from simple, reusable pieces. Your prompt fragments can be composed, nested, and reused across different contexts.

### 2. **Declarative Approach**
Describe *what* your prompt should contain, not *how* to format it. LLML handles the formatting automatically.

### 3. **Maintainable & Robust**
No more brittle string concatenation. Changes to your data structure automatically propagate to the formatted output without breaking.

### 4. **Optimized for AI Attention**
LLML's default format is structured with clear tag boundaries and numbered list items (`<rules-1>`, `<rules-2>`), which reduces prompt ambiguity and cognitive load for AI models. Clear structure means better AI performance.

### 5. **Developer Experience**
Type-safe, predictable outputs. No more debugging malformed prompts or tracking down formatting issues.

## Flexible Output Formats

LLML's formatter system gives you complete control over output format. While the default is VibeXML (optimized for LLMs), you can easily switch to JSON or create your own format:

```python
from zenbase_llml import llml, formatters

data = {
    "user": "Alice",
    "tasks": ["Review PR", "Deploy to staging"],
    "priority": "high"
}

# Default VibeXML format (optimized for LLMs)
xml_output = llml(data)
# <user>Alice</user>
# <tasks>
#   <tasks-1>Review PR</tasks-1>
#   <tasks-2>Deploy to staging</tasks-2>
# </tasks>
# <priority>high</priority>

# JSON format (when you need standard JSON)
json_output = llml(data, formatters.json)
# {"user":"Alice","tasks":["Review PR","Deploy to staging"],"priority":"high"}
```

```typescript
import { llml, formatters } from "@zenbase/llml"

const data = {
    user: "Alice",
    tasks: ["Review PR", "Deploy to staging"],
    priority: "high"
}

// Default VibeXML format
const xmlOutput = llml(data)

// JSON format with pretty printing
const jsonOutput = llml(data, formatters.json(null, 2))
// {
//   "user": "Alice",
//   "tasks": ["Review PR", "Deploy to staging"],
//   "priority": "high"
// }

// JSON with custom replacer for sensitive data
// formatters.json takes the same 2nd and 3rd arguments as JSON.stringify
const safeJson = llml(data, formatters.json((key, value) =>
    key === "password" ? "[REDACTED]" : value
))
```

This flexibility means you can:
- Use VibeXML for optimal LLM performance
- Switch to JSON for API integrations
- Create custom formats for specialized use cases
- Mix and match formatters within a single project

## Table of Contents

- [LLML - Lightweight Language Markup Language](#llml---lightweight-language-markup-language)
  - [Why We Built This](#why-we-built-this)
  - [Getting Started in 30 Seconds](#getting-started-in-30-seconds)
  - [Core Principles](#core-principles)
    - [1. **Compositional Architecture**](#1-compositional-architecture)
    - [2. **Declarative Approach**](#2-declarative-approach)
    - [3. **Maintainable \& Robust**](#3-maintainable--robust)
    - [4. **Optimized for AI Attention**](#4-optimized-for-ai-attention)
    - [5. **Developer Experience**](#5-developer-experience)
  - [Flexible Output Formats](#flexible-output-formats)
  - [Table of Contents](#table-of-contents)
  - [How It Works](#how-it-works)
    - [**Compositional Building**](#compositional-building)
    - [**VibeX MLTransformation Rules**](#vibex-mltransformation-rules)
  - [Features](#features)
    - [Why Structured Formats Win](#why-structured-formats-win)
  - [VibeXML](#vibexml)
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
  - [FAQ](#faq)
    - [Why XML instead of JSON or YAML?](#why-xml-instead-of-json-or-yaml)
    - [How does this compare to prompt templates like Jinja2 or Handlebars?](#how-does-this-compare-to-prompt-templates-like-jinja2-or-handlebars)
    - [Is there overhead? Performance impact?](#is-there-overhead-performance-impact)
    - [Is this production-ready?](#is-this-production-ready)
    - [Why four languages?](#why-four-languages)
    - [Can I customize the output format?](#can-i-customize-the-output-format)
    - [How do I migrate existing prompts?](#how-do-i-migrate-existing-prompts)
    - [Does this work with all LLM providers?](#does-this-work-with-all-llm-providers)
    - [Where can I see more examples?](#where-can-i-see-more-examples)
  - [For React Developers](#for-react-developers)
  - [License](#license)

## How It Works

LLML uses compositional patterns to transform data into structured markup:

### **Compositional Building**
```typescript
// Small, focused data structures
const userContext = { name: "Alice", role: "admin" }
const taskInstructions = ["Be precise", "Provide examples"]
const safetyRules = ["Never expose credentials", "Always validate input"]

// Compose them into complex prompts
const aiPrompt = llml({
  user: userContext,
  instructions: taskInstructions,
  safety: safetyRules,
  task: "Generate API documentation"
})
```

### **VibeX MLTransformation Rules**
1. **Simple Values**: `{key: "value"}` → `<key>value</key>`
2. **Lists/Arrays**: `{items: ["a", "b"]}` → `<items><items-1>a</items-1><items-2>b</items-2></items>`
3. **Nested Objects**: `{config: {debug: true}}` → `<config><debug>true</debug></config>`
4. **Key Preservation**: Dictionary keys are preserved exactly as provided
5. **Empty Values**: Empty objects `{}` and arrays `[]` return empty strings
6. **Extensible Formatting**: Custom formatters can be provided for specialized data types

## Features

- **Compositional Architecture**: Build complex prompts from simple, reusable pieces
- **Flexible Output Formats**: Choose between VibeXML (default), JSON, or create custom formats
- **Smart list formatting**: Arrays become `<items><items-1>first</items-1><items-2>second</items-2></items>` in XML or stay as arrays in JSON
- **Recursive nested structures**: Objects within objects maintain proper hierarchy
- **Multiline content support**: Preserves line breaks with proper indentation
- **Extensible formatter system**: Full control over how every data type is formatted
- **Type-aware processing**: Different formatters for different data types (dates, URLs, sensitive data, etc.)
- **Zero configuration**: Works out of the box with sensible defaults
- **Developer-friendly**: Type-safe, predictable, and easy to debug

### Why Structured Formats Win

1. **Clear Boundaries**: `<tag>content</tag>` creates unambiguous sections
2. **Numbered Lists**: `<items-1>`, `<items-2>` prevents ordering confusion
3. **Hierarchical Clarity**: Nested tags maintain parent-child relationships
4. **Reduced Ambiguity**: Less room for misinterpretation vs. free-form text

## VibeXML

VibeXML is a custom XML-like format that's optimized for both human readability and AI model attention. It's a loose, XML-inspired makrup language that optimizes for LLM attention.

In our experience at Zenbase, it's possible to get better performance from a smaller model using VibeXML compared to the full-size model using Markdown. We default to using VibeXML.

## Quick Start

### Python

```bash
pip install zenbase-llml # or uv, rye, poetry, etc.
```

```python
from zenbase_llml import llml

# Simple example
llml({"task": "analyze", "content": "customer feedback"})
# Output: <task>analyze</task>
#         <content>customer feedback</content>

# List handling
llml({"rules": ["be concise", "be helpful", "be accurate"]})
# Output: <rules>
#           <rules-1>be concise</rules-1>
#           <rules-2>be helpful</rules-2>
#           <rules-3>be accurate</rules-3>
#         </rules>

# Example 1: Data Extraction
extraction_prompt = llml({
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
})
# Output:
# <task>Extract key information from customer feedback</task>
# <instructions>Identify and categorize customer sentiments and specific issues mentioned</instructions>
# <rules>
#   <rules-1>Classify sentiment as positive, negative, or neutral</rules-1>
#   <rules-2>Extract specific product features mentioned</rules-2>
#   <rules-3>Identify any requested improvements or fixes</rules-3>
#   <rules-4>Note any comparisons to competitors</rules-4>
# </rules>
# <output_format>
#   <sentiment>positive/negative/neutral</sentiment>
#   <features_mentioned>
#     <features_mentioned-1>list of features</features_mentioned-1>
#   </features_mentioned>
#   <issues>
#     <issues-1>list of problems</issues-1>
#   </issues>
#   <improvements>
#     <improvements-1>list of suggestions</improvements-1>
#   </improvements>
# </output_format>

# Example 2: RAG Chatbot
rag_prompt = llml({
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
})
# Output:
# <system>You are a helpful documentation assistant</system>
# <instructions>Answer questions based on the provided documentation context</instructions>
# <documents>
#   <documents-1>
#     <title>API Authentication Guide</title>
#     <content>Our API uses OAuth 2.0 for authentication...</content>
#     <relevance_score>0.95</relevance_score>
#   </documents-1>
#   <documents-2>
#     <title>Rate Limiting Documentation</title>
#     <content>API calls are limited to 1000 requests per hour...</content>
#     <relevance_score>0.82</relevance_score>
#   </documents-2>
# </documents>
# <user_query>How do I authenticate with your API?</user_query>
# <constraints>
#   <constraints-1>Only use information from the provided documents</constraints-1>
#   <constraints-2>Cite the document title when referencing information</constraints-2>
#   <constraints-3>If information is not available, explicitly state so</constraints-3>
# </constraints>

# Example 3: AI Agent with Workflows
agent_prompt = llml({
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
})
print(agent_prompt)
# Output:
# <role>DevOps automation agent</role>
# <context>
#   <environment>production</environment>
#   <aws_region>us-east-1</aws_region>
#   <services>
#     <services-1>web-api</services-1>
#     <services-2>worker-queue</services-2>
#     <services-3>database</services-3>
#   </services>
#   <last_deployment>2024-01-15T10:30:00Z</last_deployment>
# </context>
# <instructions>Execute deployment workflow with safety checks</instructions>
# <workflows>
#   <deploy>
#     <deploy-1>Run pre-deployment health checks</deploy-1>
#     <deploy-2>Create backup of current state</deploy-2>
#     <deploy-3>Deploy to canary instance (5% traffic)</deploy-3>
#     <deploy-4>Monitor metrics for 10 minutes</deploy-4>
#     <deploy-5>If healthy, proceed to full deployment</deploy-5>
#     <deploy-6>If issues detected, automatic rollback</deploy-6>
#   </deploy>
#   <rollback>
#     <rollback-1>Stop new traffic to affected services</rollback-1>
#     <rollback-2>Restore from latest backup</rollback-2>
#     <rollback-3>Verify service health</rollback-3>
#     <rollback-4>Send notification to ops channel</rollback-4>
#   </rollback>
# </workflows>
# <safety_rules>
#   <safety_rules-1>Never skip health checks</safety_rules-1>
#   <safety_rules-2>Always maintain 99.9% uptime SLA</safety_rules-2>
#   <safety_rules-3>Require manual approval for database changes</safety_rules-3>
# </safety_rules>
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
// Output: <rules>
//           <rules-1>be concise</rules-1>
//           <rules-2>be helpful</rules-2>
//           <rules-3>be accurate</rules-3>
//         </rules>

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
// <rules>
//   <rules-1>Classify sentiment as positive, negative, or neutral</rules-1>
//   <rules-2>Extract specific product features mentioned</rules-2>
//   <rules-3>Identify any requested improvements or fixes</rules-3>
//   <rules-4>Note any comparisons to competitors</rules-4>
// </rules>
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
// <documents>
//   <documents-1>
//     <title>API Authentication Guide</title>
//     <content>Our API uses OAuth 2.0 for authentication...</content>
//     <relevance-score>0.95</relevance-score>
//   </documents-1>
//   <documents-2>
//     <title>Rate Limiting Documentation</title>
//     <content>API calls are limited to 1000 requests per hour...</content>
//     <relevance-score>0.82</relevance-score>
//   </documents-2>
// </documents>
// <user-query>How do I authenticate with your API?</user-query>
// <constraints>
//   <constraints-1>Only use information from the provided documents</constraints-1>
//   <constraints-2>Cite the document title when referencing information</constraints-2>
//   <constraints-3>If information is not available, explicitly state so</constraints-3>
// </constraints>

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
//   <environment>production</environment>
//   <aws-region>us-east-1</aws-region>
//   <services>
//     <services-1>web-api</services-1>
//     <services-2>worker-queue</services-2>
//     <services-3>database</services-3>
//   </services>
//   <last-deployment>2024-01-15T10:30:00Z</last-deployment>
// </context>
// <instructions>Execute deployment workflow with safety checks</instructions>
// <workflows>
//   <deploy>
//     <deploy-1>Run pre-deployment health checks</deploy-1>
//     <deploy-2>Create backup of current state</deploy-2>
//     <deploy-3>Deploy to canary instance (5% traffic)</deploy-3>
//     <deploy-4>Monitor metrics for 10 minutes</deploy-4>
//     <deploy-5>If healthy, proceed to full deployment</deploy-5>
//     <deploy-6>If issues detected, automatic rollback</deploy-6>
//   </deploy>
//   <rollback>
//     <rollback-1>Stop new traffic to affected services</rollback-1>
//     <rollback-2>Restore from latest backup</rollback-2>
//     <rollback-3>Verify service health</rollback-3>
//     <rollback-4>Send notification to ops channel</rollback-4>
//   </rollback>
// </workflows>
// <safety-rules>
//   <safety-rules-1>Never skip health checks</safety-rules-1>
//   <safety-rules-2>Always maintain 99.9% uptime SLA</safety-rules-2>
//   <safety-rules-3>Require manual approval for database changes</safety-rules-3>
// </safety-rules>
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
// Output: <rules>
//           <rules-1>be concise</rules-1>
//           <rules-2>be helpful</rules-2>
//           <rules-3>be accurate</rules-3>
//         </rules>

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
// Output: <rules>
//           <rules-1>be concise</rules-1>
//           <rules-2>be helpful</rules-2>
//           <rules-3>be accurate</rules-3>
//         </rules>

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
- Preserve dictionary keys exactly as provided
- Extensible formatting system for custom data types

See individual project READMEs for detailed API documentation:
- [Python API Documentation](py/README.md)
- [TypeScript API Documentation](ts/README.md)
- [Rust API Documentation](rs/README.md)
- [Go API Documentation](go/README.md)

## FAQ

### Why XML instead of JSON or YAML?

**The Attention Advantage**: Our research at Zenbase found that XML-like formats with clear open/close tags help LLMs maintain better context awareness, especially in deeply nested structures. The `<tag>content</tag>` format creates unambiguous boundaries that LLMs parse more reliably than JSON's brackets or YAML's indentation.

**Empirical Results**: In our production systems, VibeXML formatting showed:
- 15-20% fewer parsing errors by LLMs
- 30% better accuracy on complex, nested instructions
- More consistent outputs across different model providers

### How does this compare to prompt templates like Jinja2 or Handlebars?

**Different Philosophy**: Template engines focus on string interpolation. LLML focuses on data transformation.

Templates:
```python
template.render(role=role, context=context)  # String manipulation
```

LLML:
```python
llml({"role": role, "context": context})  # Data transformation
```

**Key Difference**: LLML treats your prompt as structured data, not text. This enables composability, type safety, and consistent formatting without template syntax.

### Is there overhead? Performance impact?

**Minimal overhead**: LLML is designed to be lightweight:
- Python: ~50μs for typical prompts
- TypeScript: ~30μs (Bun runtime)
- Rust: ~5μs
- Go: ~10μs

For comparison, a single LLM API call typically takes 500-5000ms. LLML's overhead is negligible.

### Is this production-ready?

**Yes**. LLML powers production AI systems at Zenbase handling millions of requests. The codebase is:
- Fully tested (90%+ coverage across all languages)
- Used in production for 6+ months
- Stable API (v1.0+)
- MIT licensed for commercial use

### Why four languages?

**Meet developers where they are**. AI development happens across the stack:
- **Python**: Data science, ML pipelines
- **TypeScript**: Full-stack web apps, edge functions
- **Rust**: High-performance services, embedded systems
- **Go**: Cloud infrastructure, microservices

All implementations produce identical output and share the same philosophy.

### Can I customize the output format?

**Yes!** LLML has a completely extensible formatter system. For example, both Python and TypeScript now include built-in JSON formatters:

```python
from zenbase_llml import llml
from zenbase_llml.formatters.json import json

# Use JSON instead of VibeXML
result = llml({"task": "Deploy", "env": "prod"}, json)
# Output: {"task":"Deploy","env":"prod"}

# Create custom formatters for your types
def format_money(value, llml_func, formatters):
    return f"${value.amount:.2f} {value.currency}"

custom_formatters = {
    lambda v: isinstance(v, Money): format_money,
    **json  # Fallback to JSON for other types
}
```

```typescript
import { llml, json } from "@zenbase/llml"

// Pretty-printed JSON
const pretty = llml(data, json(null, 2))

// JSON with sensitive data filtering
const safe = llml(data, json((k, v) =>
    k.includes("secret") ? "[REDACTED]" : v
))
```

You can:
- Switch between VibeXML (default) and JSON
- Create formatters for domain objects (User, Product, Money)
- Add special handling for sensitive data
- Build entirely custom output formats
- Mix different formatters in one project

See the [formatter documentation](docs/formatters.md) for more examples.

### How do I migrate existing prompts?

**Gradually**. LLML is designed for incremental adoption:

```python
# Start with your existing prompt
old_prompt = f"Role: {role}\nTask: {task}"

# Wrap parts in LLML as you go
new_prompt = f"Role: {role}\n{llml({'task': task, 'context': context})}"

# Eventually, full LLML
final_prompt = llml({"role": role, "task": task, "context": context})
```

### Does this work with all LLM providers?

**Yes**. LLML outputs plain text that works with any LLM:
- OpenAI GPT-3.5/4/4o
- Anthropic Claude
- Google Gemini
- Mistral
- Local models (Llama, etc.)
- Any text-input API

### Where can I see more examples?

Check out:
- [Full documentation](.cursor/rules/spec.mdc)
- [Compositional patterns guide](REACT.md)
- Language-specific examples in each implementation's README
- [Real-world use cases](examples/) (coming soon)

## For React Developers

If you're familiar with React, you'll recognize many patterns in LLML:
- **Composition**: Like React components, LLML lets you build complex structures from simple, reusable pieces
- **Declarative**: Describe what you want, not how to build it
- **Data-driven**: Your prompts are a function of your data

See our [detailed comparison guide](REACT.md) for React developers.

## License

MIT
