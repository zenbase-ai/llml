package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestDirectArray(t *testing.T) {
	result := llml.Sprintf([]any{"a", "b", "c"})
	expected := "<1>a</1>\n<2>b</2>\n<3>c</3>"
	assert.Equal(t, expected, result)
}

func TestDirectArrayWithMixedTypes(t *testing.T) {
	result := llml.Sprintf([]any{1, "hello", true})
	expected := "<1>1</1>\n<2>hello</2>\n<3>true</3>"
	assert.Equal(t, expected, result)
}

func TestDirectArrayWithObjects(t *testing.T) {
	result := llml.Sprintf([]any{
		map[string]any{"name": "Alice"},
		map[string]any{"name": "Bob"},
	})
	expected := "<1>\n  <1-name>Alice</1-name>\n</1>\n<2>\n  <2-name>Bob</2-name>\n</2>"
	assert.Equal(t, expected, result)
}

func TestDirectEmptyArray(t *testing.T) {
	result := llml.Sprintf([]any{})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestDirectArrayWithIndentation(t *testing.T) {
	result := llml.Sprintf([]any{"a", "b"}, llml.Options{Indent: "  "})
	expected := "  <1>a</1>\n  <2>b</2>"
	assert.Equal(t, expected, result)
}

func TestDirectArrayWithPrefix(t *testing.T) {
	result := llml.Sprintf([]any{"a", "b"}, llml.Options{Prefix: "item"})
	expected := "<item-1>a</item-1>\n<item-2>b</item-2>"
	assert.Equal(t, expected, result)
}
