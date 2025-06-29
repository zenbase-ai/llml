# LLML - Lightweight Language Markup Language

Convert Python and JavaScript data structures into human-readable, XML-like markup.

## Overview

LLML is a data serialization library available in both Python and TypeScript/JavaScript that transforms nested data structures (dictionaries/objects, lists/arrays, primitives) into well-formatted, XML-like markup. It's designed for tasks like AI prompt engineering, configuration file generation, and creating structured documents from data.

## Features

- **Automatic kebab-case conversion** of keys
- **Smart list formatting** with numbered items
- **Recursive handling** of deeply nested structures
- **Automatic indentation** for readability
- **Multiline content support**
- **Optional prefix namespacing**
- **Zero configuration** - works out of the box

## Quick Start

### Python
```python
from llml import llml

# Simple example
print(llml(task="analyze", content="customer feedback"))
# Output: <task>analyze</task>
#         <content>customer feedback</content>

# List handling
print(llml(rules=["be concise", "be helpful", "be accurate"]))
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
print(extraction_prompt)
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
print(rag_prompt)
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
console.log(extractionPrompt);

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
console.log(ragPrompt);

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
console.log(agentPrompt);
```

### Rust
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
}), None);
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
}), None);
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
}), None);
println!("{}", agent_prompt);
```

### Go
```go
package main

import (
	"fmt"
	"github.com/zenbase-ai/llml-go/pkg/llml"
)

func main() {
	// Simple example
	fmt.Println(llml.LLML(map[string]any{
		"task": "analyze",
		"content": "customer feedback",
	}))
	// Output: <task>analyze</task>
	//         <content>customer feedback</content>

	// List handling
	fmt.Println(llml.LLML(map[string]any{
		"rules": []any{"be concise", "be helpful", "be accurate"},
	}))
	// Output: <rules-list>
	//           <rules-1>be concise</rules-1>
	//           <rules-2>be helpful</rules-2>
	//           <rules-3>be accurate</rules-3>
	//         </rules-list>

	// Example 1: Data Extraction
	extractionPrompt := llml.LLML(map[string]any{
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
	fmt.Println(extractionPrompt)

	// Example 2: RAG Chatbot
	ragPrompt := llml.LLML(map[string]any{
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
	fmt.Println(ragPrompt)

	// Example 3: AI Agent with Workflows
	agentPrompt := llml.LLML(map[string]any{
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
	fmt.Println(agentPrompt)
}
```

## Project Structure

This repository contains two implementations:

- **`py/`** - Python implementation with full test suite
- **`ts/`** - TypeScript implementation with comprehensive tests
- **`rs/`** - Rust implementation with comprehensive tests
- **`go/`** - Go implementation with comprehensive tests

Both implementations provide identical functionality and API design.

## Installation & Setup

### Python Project
```bash
cd py/
uv venv
source .venv/bin/activate
uv pip install -e .[dev]
```

### TypeScript Project
```bash
cd ts/
bun install
```

## Running Tests

### All Tests
```bash
just test
```

### Python Tests Only
```bash
just py test
```

### TypeScript Tests Only
```bash
just ts test
```

## Development

This project uses:
- **Python**: `uv` for dependency management, `pytest` for testing, `ruff` for linting
- **TypeScript**: `bun` for runtime and package management, `vitest` for testing
- **Rust**: `cargo` for dependency management, `cargo test` for testing
- **Go**: `go` for dependency management, `go test` for testing
- **Just**: Task runner for cross-platform commands

## API Reference

Both implementations provide the same core functionality:

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
