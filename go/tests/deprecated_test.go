package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Test deprecated LLML function for backwards compatibility

func TestDeprecatedLLMLFunction(t *testing.T) {
	// Test that LLML function works the same as Sprintf
	data := map[string]any{
		"name": "Alice",
		"age":  30,
	}
	
	sprintfResult := llml.Sprintf(data)
	llmlResult := llml.LLML(data)
	
	// Both should produce identical output
	assert.Equal(t, sprintfResult, llmlResult)
	
	// Verify the actual content
	expected := "<age>30</age>\n<name>Alice</name>"
	assert.Equal(t, expected, llmlResult)
}

func TestDeprecatedLLMLWithOptions(t *testing.T) {
	// Test LLML function with options
	data := map[string]any{
		"name": "Bob",
		"items": []any{"item1", "item2"},
	}
	
	opts := llml.Options{
		Indent: "  ",
		Prefix: "test",
		Strict: true,
	}
	
	sprintfResult := llml.Sprintf(data, opts)
	llmlResult := llml.LLML(data, opts)
	
	// Both should produce identical output
	assert.Equal(t, sprintfResult, llmlResult)
	
	// Verify it actually uses the options
	assert.Contains(t, llmlResult, "test-")
	assert.Contains(t, llmlResult, "  <") // Check indentation
}

func TestDeprecatedLLMLWithNil(t *testing.T) {
	// Test LLML function with nil input
	sprintfResult := llml.Sprintf(nil)
	llmlResult := llml.LLML(nil)
	
	assert.Equal(t, sprintfResult, llmlResult)
	assert.Equal(t, "nil", llmlResult)
}

func TestDeprecatedLLMLWithEmptyData(t *testing.T) {
	// Test LLML function with empty data
	data := map[string]any{}
	
	sprintfResult := llml.Sprintf(data)
	llmlResult := llml.LLML(data)
	
	assert.Equal(t, sprintfResult, llmlResult)
	assert.Equal(t, "", llmlResult)
}