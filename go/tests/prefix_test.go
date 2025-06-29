package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestWithPrefix(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"config": "value",
	}, llml.Options{Prefix: "app"})
	expected := "<app-config>value</app-config>"
	assert.Equal(t, expected, result)
}

func TestListWithPrefix(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"items": []any{"a", "b"},
	}, llml.Options{Prefix: "app"})
	expected := "<app-items>\n" +
		"  <app-items-1>a</app-items-1>\n" +
		"  <app-items-2>b</app-items-2>\n" +
		"</app-items>"
	assert.Equal(t, expected, result)
}
