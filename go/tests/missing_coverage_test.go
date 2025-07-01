package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Tests targeting specific uncovered code paths to maximize coverage

// Test direct map[string]interface{} input to Sprintf (Lines 40-46)
func TestDirectMapStringInterfaceInput(t *testing.T) {
	// This targets the uncovered path in Sprintf when called directly with map[string]interface{}
	data := map[string]interface{}{
		"name": "Direct Test",
		"value": 42,
	}
	result := llml.Sprintf(data)
	expected := "<name>Direct Test</name>\n<value>42</value>"
	assert.Equal(t, expected, result)
}

// Test direct []interface{} input to Sprintf (Lines 53-57)
func TestDirectSliceInterfaceInput(t *testing.T) {
	// This targets the uncovered path in Sprintf when called directly with []interface{}
	data := []interface{}{"direct", "slice", "test"}
	result := llml.Sprintf(data)
	expected := "<1>direct</1>\n<2>slice</2>\n<3>test</3>"
	assert.Equal(t, expected, result)
}

// Test []interface{} in formatKeyValue (Lines 139-143)
func TestSliceInterfaceInFormatKeyValue(t *testing.T) {
	// This targets the uncovered path in formatKeyValue with []interface{} slices
	data := map[string]any{
		"items": []interface{}{"a", "b", "c"},
	}
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n  <items-3>c</items-3>\n</items>"
	assert.Equal(t, expected, result)
}

// Test map[string]interface{} in formatKeyValue (Lines 149-154)
func TestMapStringInterfaceInFormatKeyValue(t *testing.T) {
	// This targets the uncovered path in formatKeyValue with map[string]interface{} nested objects
	data := map[string]any{
		"config": map[string]interface{}{
			"enabled": true,
			"timeout": 30,
		},
	}
	result := llml.Sprintf(data)
	expected := "<config>\n  <enabled>true</enabled>\n  <timeout>30</timeout>\n</config>"
	assert.Equal(t, expected, result)
}

// Test map[string]interface{} objects in formatList (Lines 227-245)
func TestMapStringInterfaceInFormatList(t *testing.T) {
	// This targets the uncovered path in formatList with map[string]interface{} objects
	data := map[string]any{
		"users": []any{
			map[string]interface{}{
				"id":   1,
				"name": "Alice",
			},
			map[string]interface{}{
				"id":   2,
				"name": "Bob",
			},
		},
	}
	result := llml.Sprintf(data)
	expected := "<users>\n  <users-1>\n    <id>1</id>\n    <name>Alice</name>\n  </users-1>\n  <users-2>\n    <id>2</id>\n    <name>Bob</name>\n  </users-2>\n</users>"
	assert.Equal(t, expected, result)
}

// Test map[string]interface{} in formatSlice (Lines 292-315)
func TestMapStringInterfaceInFormatSlice(t *testing.T) {
	// This targets the uncovered path in formatSlice with map[string]interface{} objects
	data := []any{
		map[string]interface{}{
			"type": "user",
			"id":   1,
		},
		map[string]interface{}{
			"type": "admin",
			"id":   2,
		},
	}
	result := llml.Sprintf(data)
	expected := "<1>\n  <1-id>1</1-id>\n  <1-type>user</1-type>\n</1>\n<2>\n  <2-id>2</2-id>\n  <2-type>admin</2-type>\n</2>"
	assert.Equal(t, expected, result)
}

// Test []interface{} nested arrays in formatSlice (Lines 331-346)
func TestSliceInterfaceNestedInFormatSlice(t *testing.T) {
	// This targets the uncovered path in formatSlice with []interface{} nested arrays
	data := []any{
		[]interface{}{"first", "array"},
		[]interface{}{"second", "array"},
	}
	result := llml.Sprintf(data)
	expected := "<1>\n  <1>first</1>\n  <2>array</2>\n</1>\n<2>\n  <1>second</1>\n  <2>array</2>\n</2>"
	assert.Equal(t, expected, result)
}

// Test empty map[string]interface{} in formatSlice to hit both branches
func TestEmptyMapStringInterfaceInFormatSlice(t *testing.T) {
	// This targets both the empty and non-empty branches in formatSlice for map[string]interface{}
	data := []any{
		map[string]interface{}{}, // Empty - should create <1></1>
		map[string]interface{}{   // Non-empty - should create full structure
			"name": "test",
		},
	}
	result := llml.Sprintf(data)
	expected := "<1></1>\n<2>\n  <2-name>test</2-name>\n</2>"
	assert.Equal(t, expected, result)
}

// Test empty []interface{} in formatSlice to ensure proper handling
func TestEmptySliceInterfaceInFormatSlice(t *testing.T) {
	// This tests the empty []interface{} handling in formatSlice
	data := []any{
		[]interface{}{}, // Empty - should be skipped
		[]interface{}{"not", "empty"},
	}
	result := llml.Sprintf(data)
	// Empty array at position 1 should be skipped, so only position 2 appears
	expected := "<2>\n  <1>not</1>\n  <2>empty</2>\n</2>"
	assert.Equal(t, expected, result)
}

// Test complex mixed interface types to hit multiple paths
func TestComplexMixedInterfaceTypes(t *testing.T) {
	// This combines multiple interface conversion paths in one test
	data := map[string]any{
		"directMap": map[string]interface{}{
			"nested": map[string]interface{}{
				"value": "deep",
			},
			"array": []interface{}{"item1", "item2"},
		},
		"directArray": []interface{}{
			map[string]interface{}{
				"id": 1,
			},
			[]interface{}{"nested", "array"},
		},
	}
	result := llml.Sprintf(data)
	
	// Verify key components are present
	assert.Contains(t, result, "<directArray>")
	assert.Contains(t, result, "<directMap>")
	assert.Contains(t, result, "<nested>")
	assert.Contains(t, result, "<value>deep</value>")
	assert.Contains(t, result, "<array>")
	assert.Contains(t, result, "<array-1>item1</array-1>")
}