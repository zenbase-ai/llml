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
	assert.Contains(t, result, "<user-config>")
	assert.Contains(t, result, "</user-config>")
	// In non-strict mode (default), nested object properties are kebab-cased but don't include parent prefixes
	assert.Contains(t, result, "<debug-mode>true</debug-mode>")
	assert.Contains(t, result, "<max-retries>5</max-retries>")
}

func TestNestedCamelCaseInObjects(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"userConfig": map[string]any{
			"debugMode":  true,
			"maxRetries": 5,
			"XMLParser":  "enabled",
		},
	})
	assert.Contains(t, result, "<user-config>")
	assert.Contains(t, result, "</user-config>")
	// In non-strict mode (default), nested object properties are kebab-cased but don't include parent prefixes
	assert.Contains(t, result, "<debug-mode>true</debug-mode>")
	assert.Contains(t, result, "<max-retries>5</max-retries>")
	assert.Contains(t, result, "<xml-parser>enabled</xml-parser>")
}
