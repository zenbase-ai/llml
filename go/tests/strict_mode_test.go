package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestStrictModeFalseNestedObjects(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"config": map[string]any{
			"debug":   true,
			"timeout": 30,
		},
	}, llml.Options{Strict: false})
	// In non-strict mode, nested object properties don't include parent key prefixes
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "</config>")
	assert.Contains(t, result, "<debug>true</debug>")
	assert.Contains(t, result, "<timeout>30</timeout>")
}

func TestStrictModeTrueNestedObjects(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"config": map[string]any{
			"debug":   true,
			"timeout": 30,
		},
	}, llml.Options{Strict: true})
	// In strict mode, nested object properties include parent key prefixes
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "</config>")
	assert.Contains(t, result, "<config-debug>true</config-debug>")
	assert.Contains(t, result, "<config-timeout>30</config-timeout>")
}

func TestStrictModeFalseArrayObjects(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"data": []any{
			map[string]any{"name": "Alice", "age": 30},
			map[string]any{"name": "Bob", "age": 25},
		},
	}, llml.Options{Strict: false})
	// In non-strict mode, object properties within arrays don't include array item prefixes
	assert.Contains(t, result, "<data>")
	assert.Contains(t, result, "</data>")
	assert.Contains(t, result, "<data-1>")
	assert.Contains(t, result, "</data-1>")
	assert.Contains(t, result, "<data-2>")
	assert.Contains(t, result, "</data-2>")
	assert.Contains(t, result, "<name>Alice</name>")
	assert.Contains(t, result, "<age>30</age>")
	assert.Contains(t, result, "<name>Bob</name>")
	assert.Contains(t, result, "<age>25</age>")
}

func TestStrictModeTrueArrayObjects(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"data": []any{
			map[string]any{"name": "Alice", "age": 30},
			map[string]any{"name": "Bob", "age": 25},
		},
	}, llml.Options{Strict: true})
	// In strict mode, object properties within arrays include array item prefixes
	assert.Contains(t, result, "<data>")
	assert.Contains(t, result, "</data>")
	assert.Contains(t, result, "<data-1>")
	assert.Contains(t, result, "</data-1>")
	assert.Contains(t, result, "<data-2>")
	assert.Contains(t, result, "</data-2>")
	assert.Contains(t, result, "<data-1-name>Alice</data-1-name>")
	assert.Contains(t, result, "<data-1-age>30</data-1-age>")
	assert.Contains(t, result, "<data-2-name>Bob</data-2-name>")
	assert.Contains(t, result, "<data-2-age>25</data-2-age>")
}

func TestStrictModeFalseWithKebabCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user_config": map[string]any{
			"debug_mode": true,
			"maxRetries": 5,
		},
	}, llml.Options{Strict: false})
	assert.Contains(t, result, "<user-config>")
	assert.Contains(t, result, "</user-config>")
	// In non-strict mode, nested properties are kebab-cased but don't include parent prefixes
	assert.Contains(t, result, "<debug-mode>true</debug-mode>")
	assert.Contains(t, result, "<max-retries>5</max-retries>")
}

func TestStrictModeTrueWithKebabCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user_config": map[string]any{
			"debug_mode": true,
			"maxRetries": 5,
		},
	}, llml.Options{Strict: true})
	assert.Contains(t, result, "<user-config>")
	assert.Contains(t, result, "</user-config>")
	// In strict mode, nested properties are kebab-cased and include parent prefixes
	assert.Contains(t, result, "<user-config-debug-mode>true</user-config-debug-mode>")
	assert.Contains(t, result, "<user-config-max-retries>5</user-config-max-retries>")
}

func TestDefaultStrictModeBehavior(t *testing.T) {
	// Default behavior should be strict: false
	result := llml.Sprintf(map[string]any{
		"config": map[string]any{
			"debug": true,
		},
	})
	// Should behave like non-strict mode by default
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "<debug>true</debug>")
	assert.NotContains(t, result, "<config-debug>")
}