package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestNestedDict(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"config": map[string]any{
			"debug":   true,
			"timeout": 30,
		},
	})
	// In non-strict mode (default), nested object properties don't include parent key prefixes
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "</config>")
	assert.Contains(t, result, "<debug>true</debug>")
	assert.Contains(t, result, "<timeout>30</timeout>")
}

func TestNestedDictWithKebabCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user_config": map[string]any{
			"debug_mode": true,
			"maxRetries": 5,
		},
	})
	assert.Contains(t, result, "<user_config>")
	assert.Contains(t, result, "</user_config>")
	// In non-strict mode (default), nested object properties use original key names
	assert.Contains(t, result, "<debug_mode>true</debug_mode>")
	assert.Contains(t, result, "<maxRetries>5</maxRetries>")
}

func TestNestedCamelCaseInObjects(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"userConfig": map[string]any{
			"debugMode":  true,
			"maxRetries": 5,
			"XMLParser":  "enabled",
		},
	})
	assert.Contains(t, result, "<userConfig>")
	assert.Contains(t, result, "</userConfig>")
	// In non-strict mode (default), nested object properties use original key names
	assert.Contains(t, result, "<debugMode>true</debugMode>")
	assert.Contains(t, result, "<maxRetries>5</maxRetries>")
	assert.Contains(t, result, "<XMLParser>enabled</XMLParser>")
}
