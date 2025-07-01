package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Integration tests combining multiple options and complex scenarios

func TestIndentPrefixStrictCombination(t *testing.T) {
	// Test all options combined: Indent + Prefix + Strict
	data := map[string]any{
		"user": map[string]any{
			"name": "Alice",
			"profile": map[string]any{
				"age":  30,
				"city": "New York",
			},
		},
		"items": []any{
			map[string]any{
				"id":   1,
				"name": "Item 1",
			},
		},
	}
	
	opts := llml.Options{
		Indent: "    ", // 4 spaces
		Prefix: "api",
		Strict: true,
	}
	
	result := llml.Sprintf(data, opts)
	
	// Verify it includes all options
	assert.Contains(t, result, "api-") // Prefix
	assert.Contains(t, result, "    <") // Indentation
	assert.Contains(t, result, "api-user-") // Strict mode prefix propagation
}

func TestComplexRealWorldExample(t *testing.T) {
	// Test a realistic complex data structure
	data := map[string]any{
		"apiResponse": map[string]any{
			"status": "success",
			"data": map[string]any{
				"users": []any{
					map[string]any{
						"id":    1,
						"name":  "Alice Johnson",
						"email": "alice@example.com",
						"profile": map[string]any{
							"age":         30,
							"city":        "New York",
							"preferences": []any{"music", "reading", "travel"},
						},
					},
					map[string]any{
						"id":    2,
						"name":  "Bob Smith",
						"email": "bob@example.com",
						"profile": map[string]any{
							"age":         25,
							"city":        "Los Angeles",
							"preferences": []any{"sports", "gaming"},
						},
					},
				},
				"pagination": map[string]any{
					"page":  1,
					"limit": 10,
					"total": 2,
				},
			},
			"timestamp": "2024-01-01T12:00:00Z",
		},
	}
	
	result := llml.Sprintf(data)
	
	// Verify structure exists and is properly formatted
	assert.Contains(t, result, "<apiResponse>")
	assert.Contains(t, result, "<users>")
	assert.Contains(t, result, "<users-1>")
	assert.Contains(t, result, "<users-2>")
	assert.Contains(t, result, "<preferences>")
	assert.Contains(t, result, "<preferences-1>music</preferences-1>")
	assert.Contains(t, result, "<pagination>")
	assert.Contains(t, result, "<timestamp>2024-01-01T12:00:00Z</timestamp>")
}

func TestStrictModeVsNonStrictMode(t *testing.T) {
	// Test difference between strict and non-strict modes
	data := map[string]any{
		"container": map[string]any{
			"nested": map[string]any{
				"value": "test",
			},
		},
	}
	
	// Non-strict mode (default)
	nonStrictResult := llml.Sprintf(data)
	
	// Strict mode
	strictOpts := llml.Options{Strict: true}
	strictResult := llml.Sprintf(data, strictOpts)
	
	// In non-strict mode, nested keys should not have prefix
	assert.Contains(t, nonStrictResult, "<value>test</value>")
	
	// In strict mode, nested keys should have prefix
	assert.Contains(t, strictResult, "<container-nested-value>test</container-nested-value>")
}

func TestArraysInStrictMode(t *testing.T) {
	// Test how arrays behave differently in strict vs non-strict mode
	data := map[string]any{
		"items": []any{
			map[string]any{
				"name":  "Item 1",
				"value": 100,
			},
			map[string]any{
				"name":  "Item 2", 
				"value": 200,
			},
		},
	}
	
	// Non-strict mode
	nonStrictResult := llml.Sprintf(data)
	
	// Strict mode
	strictOpts := llml.Options{Strict: true}
	strictResult := llml.Sprintf(data, strictOpts)
	
	// In non-strict mode, array item contents don't have prefix
	assert.Contains(t, nonStrictResult, "<name>Item 1</name>")
	
	// In strict mode, array item contents have prefix
	assert.Contains(t, strictResult, "<items-1-name>Item 1</items-1-name>")
}

func TestPrefixWithComplexKeys(t *testing.T) {
	// Test prefix behavior with complex key names
	data := map[string]any{
		"userInfo": map[string]any{
			"firstName": "John",
			"lastName":  "Doe",
		},
		"itemList": []any{
			map[string]any{
				"itemName": "Product A",
				"itemType": "electronics",
			},
		},
	}
	
	opts := llml.Options{
		Prefix: "api",
		Strict: true,
	}
	
	result := llml.Sprintf(data, opts)
	
	// Verify prefix is applied with original key names
	assert.Contains(t, result, "<api-userInfo>")
	assert.Contains(t, result, "<api-userInfo-firstName>John</api-userInfo-firstName>")
	assert.Contains(t, result, "<api-itemList>")
	assert.Contains(t, result, "<api-itemList-1-itemName>Product A</api-itemList-1-itemName>")
}

func TestMixedDataTypesWithOptions(t *testing.T) {
	// Test mixed data types with various options
	data := map[string]any{
		"stringValue": "hello world",
		"intValue":    42,
		"floatValue":  3.14159,
		"boolValue":   true,
		"nilValue":    nil,
		"arrayValue":  []any{1, 2, 3},
		"objectValue": map[string]any{
			"nested": "value",
		},
		"multilineValue": "Line 1\nLine 2\nLine 3",
	}
	
	opts := llml.Options{
		Indent: "  ",
		Prefix: "test",
		Strict: false,
	}
	
	result := llml.Sprintf(data, opts)
	
	// Verify all types are handled correctly with options
	assert.Contains(t, result, "<test-stringValue>hello world</test-stringValue>")
	assert.Contains(t, result, "<test-intValue>42</test-intValue>")
	assert.Contains(t, result, "<test-floatValue>3.14159</test-floatValue>")
	assert.Contains(t, result, "<test-boolValue>true</test-boolValue>")
	assert.Contains(t, result, "<test-nilValue>nil</test-nilValue>")
	assert.Contains(t, result, "<test-arrayValue>")
	assert.Contains(t, result, "<test-objectValue>")
	// Go implementation doesn't add extra indentation to multiline content
	assert.Contains(t, result, "<test-multilineValue>\n  Line 1\n  Line 2\n  Line 3\n  </test-multilineValue>")
}