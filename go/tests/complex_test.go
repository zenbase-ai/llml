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
	assert.Contains(t, result, "<sections-list>")
	assert.Contains(t, result, "<sections-1>intro</sections-1>")
	assert.Contains(t, result, "<sections-2>body</sections-2>")
	assert.Contains(t, result, "<sections-3>conclusion</sections-3>")
	assert.Contains(t, result, "</sections-list>")
	assert.Contains(t, result, "<metadata>")
	assert.Contains(t, result, "<metadata-author>Alice</metadata-author>")
	assert.Contains(t, result, "<metadata-version>1.0</metadata-version>")
	assert.Contains(t, result, "</metadata>")
}
