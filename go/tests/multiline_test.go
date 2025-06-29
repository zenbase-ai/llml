package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestMultilineContent(t *testing.T) {
	content := `
    Line 1
    Line 2
    Line 3
    `
	result := llml.Sprintf(map[string]any{
		"description": content,
	})
	expected := "<description>\n  Line 1\n  Line 2\n  Line 3\n</description>"
	assert.Equal(t, expected, result)
}
