package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestBasicIndentation(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"message": "Hello",
	}, llml.Options{Indent: "  "})
	expected := "  <message>Hello</message>"
	assert.Equal(t, expected, result)
}

func TestListWithIndentation(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"items": []any{"a", "b"},
	}, llml.Options{Indent: "  "})
	expected := "  <items-list>\n" +
		"    <items-1>a</items-1>\n" +
		"    <items-2>b</items-2>\n" +
		"  </items-list>"
	assert.Equal(t, expected, result)
}
