package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestMixedContent(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"title":    "My Document",
		"sections": []any{"intro", "body", "conclusion"},
		"metadata": map[string]any{
			"author":  "Alice",
			"version": "1.0",
		},
	})
	// Check for key components
	assert.Contains(t, result, "<title>My Document</title>")
	assert.Contains(t, result, "<sections>")
	assert.Contains(t, result, "<sections-1>intro</sections-1>")
	assert.Contains(t, result, "<sections-2>body</sections-2>")
	assert.Contains(t, result, "<sections-3>conclusion</sections-3>")
	assert.Contains(t, result, "</sections>")
	assert.Contains(t, result, "<metadata>")
	// In non-strict mode (default), nested object properties don't include parent prefixes
	assert.Contains(t, result, "<author>Alice</author>")
	assert.Contains(t, result, "<version>1.0</version>")
	assert.Contains(t, result, "</metadata>")
}
