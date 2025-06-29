package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Note: Direct array tests - Go implementation returns empty string for direct arrays
// This matches the SPEC.md requirement for direct array handling
func TestDirectArrayReturnsEmpty(t *testing.T) {
	result := llml.Sprintf([]any{"a", "b", "c"})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestDirectArrayWithMixedTypesReturnsEmpty(t *testing.T) {
	result := llml.Sprintf([]any{1, "hello", true})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestDirectArrayWithObjectsReturnsEmpty(t *testing.T) {
	result := llml.Sprintf([]any{
		map[string]any{"name": "Alice"},
		map[string]any{"name": "Bob"},
	})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestDirectEmptyArrayReturnsEmpty(t *testing.T) {
	result := llml.Sprintf([]any{})
	expected := ""
	assert.Equal(t, expected, result)
}
