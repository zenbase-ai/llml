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
	// Check that both nested elements are present
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "</config>")
	assert.Contains(t, result, "<config-debug>true</config-debug>")
	assert.Contains(t, result, "<config-timeout>30</config-timeout>")
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
	assert.Contains(t, result, "<user-config-debug-mode>true</user-config-debug-mode>")
	assert.Contains(t, result, "<user-config-max-retries>5</user-config-max-retries>")
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
	assert.Contains(t, result, "<user-config-debug-mode>true</user-config-debug-mode>")
	assert.Contains(t, result, "<user-config-max-retries>5</user-config-max-retries>")
	assert.Contains(t, result, "<user-config-xml-parser>enabled</user-config-xml-parser>")
}
