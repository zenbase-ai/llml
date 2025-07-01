package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Tests designed to hit the exact remaining uncovered lines to reach 90%+ coverage

// Force interface{} types at runtime to hit actual conversion paths

// Test Lines 40-46: Direct map[string]interface{} input to Sprintf
func TestRealMapStringInterfaceDirectInput(t *testing.T) {
	// Create a function that returns map[string]interface{} to force the type
	getData := func() map[string]interface{} {
		return map[string]interface{}{
			"key1": "value1",
			"key2": 42,
		}
	}
	
	// Use interface{} variable to force type checking
	var data interface{} = getData()
	result := llml.Sprintf(data)
	expected := "<key1>value1</key1>\n<key2>42</key2>"
	assert.Equal(t, expected, result)
}

// Test Lines 53-58: Direct []interface{} input to Sprintf  
func TestRealSliceInterfaceDirectInput(t *testing.T) {
	// Create a function that returns []interface{} to force the type
	getData := func() []interface{} {
		return []interface{}{"item1", "item2", 123}
	}
	
	// Use interface{} variable to force type checking
	var data interface{} = getData()
	result := llml.Sprintf(data)
	expected := "<1>item1</1>\n<2>item2</2>\n<3>123</3>"
	assert.Equal(t, expected, result)
}

// Test Lines 139-143: []interface{} in formatKeyValue
func TestRealSliceInterfaceInKeyValue(t *testing.T) {
	// Force []interface{} type as a value in a map
	getSlice := func() interface{} {
		return []interface{}{"a", "b", "c"}
	}
	
	data := map[string]any{
		"items": getSlice(),
	}
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n  <items-3>c</items-3>\n</items>"
	assert.Equal(t, expected, result)
}

// Test Lines 149-154: map[string]interface{} in formatKeyValue
func TestRealMapStringInterfaceInKeyValue(t *testing.T) {
	// Force map[string]interface{} type as a value in a map
	getMap := func() interface{} {
		return map[string]interface{}{
			"nested": "value",
			"count":  5,
		}
	}
	
	data := map[string]any{
		"config": getMap(),
	}
	result := llml.Sprintf(data)
	expected := "<config>\n  <count>5</count>\n  <nested>value</nested>\n</config>"
	assert.Equal(t, expected, result)
}

// Test Lines 227-245: map[string]interface{} objects in formatList
func TestRealMapStringInterfaceInArray(t *testing.T) {
	// Force map[string]interface{} objects in an array
	getMapInterface := func() interface{} {
		return map[string]interface{}{
			"id":   1,
			"name": "test",
		}
	}
	
	data := map[string]any{
		"items": []any{
			getMapInterface(),
			getMapInterface(),
		},
	}
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>\n    <id>1</id>\n    <name>test</name>\n  </items-1>\n  <items-2>\n    <id>1</id>\n    <name>test</name>\n  </items-2>\n</items>"
	assert.Equal(t, expected, result)
}

// Test Lines 292-315: map[string]interface{} in formatSlice (direct array)
func TestRealMapStringInterfaceInDirectArray(t *testing.T) {
	// Force map[string]interface{} in direct array
	getMapInterface := func() interface{} {
		return map[string]interface{}{
			"type": "user",
			"id":   42,
		}
	}
	
	data := []any{
		getMapInterface(),
		getMapInterface(),
	}
	result := llml.Sprintf(data)
	expected := "<1>\n  <1-id>42</1-id>\n  <1-type>user</1-type>\n</1>\n<2>\n  <2-id>42</2-id>\n  <2-type>user</2-type>\n</2>"
	assert.Equal(t, expected, result)
}

// Test Lines 331-345: []interface{} nested arrays in formatSlice
func TestRealSliceInterfaceInDirectArray(t *testing.T) {
	// Force []interface{} nested in direct array
	getSliceInterface := func() interface{} {
		return []interface{}{"nested", "items"}
	}
	
	data := []any{
		getSliceInterface(),
		getSliceInterface(),
	}
	result := llml.Sprintf(data)
	expected := "<1>\n  <1>nested</1>\n  <2>items</2>\n</1>\n<2>\n  <1>nested</1>\n  <2>items</2>\n</2>"
	assert.Equal(t, expected, result)
}

// Test empty map[string]interface{} to hit both branches in formatSlice
func TestEmptyMapStringInterfaceInDirectArrayForced(t *testing.T) {
	getEmptyMapInterface := func() interface{} {
		return map[string]interface{}{}
	}
	
	getNonEmptyMapInterface := func() interface{} {
		return map[string]interface{}{
			"key": "value",
		}
	}
	
	data := []any{
		getEmptyMapInterface(),
		getNonEmptyMapInterface(),
	}
	result := llml.Sprintf(data)
	expected := "<1></1>\n<2>\n  <2-key>value</2-key>\n</2>"
	assert.Equal(t, expected, result)
}

// Test empty []interface{} in direct array
func TestEmptySliceInterfaceInDirectArrayForced(t *testing.T) {
	getEmptySliceInterface := func() interface{} {
		return []interface{}{}
	}
	
	getNonEmptySliceInterface := func() interface{} {
		return []interface{}{"item"}
	}
	
	data := []any{
		getEmptySliceInterface(), // Should be skipped
		getNonEmptySliceInterface(),
	}
	result := llml.Sprintf(data)
	// Empty slice should be skipped, so only item 2 appears
	expected := "<2>\n  <1>item</1>\n</2>"
	assert.Equal(t, expected, result)
}

// Complex test that combines multiple interface{} types to hit various paths
func TestComplexInterfaceTypeForcing(t *testing.T) {
	// Build complex structure with forced interface{} types
	getComplexData := func() interface{} {
		return map[string]interface{}{
			"users": []interface{}{
				map[string]interface{}{
					"id":   1,
					"tags": []interface{}{"admin", "active"},
				},
				map[string]interface{}{
					"id":   2,
					"tags": []interface{}{"user"},
				},
			},
			"config": map[string]interface{}{
				"timeout": 30,
				"enabled": true,
			},
		}
	}
	
	var data interface{} = getComplexData()
	result := llml.Sprintf(data)
	
	// Verify structure exists
	assert.Contains(t, result, "<users>")
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "<users-1>")
	assert.Contains(t, result, "<tags>")
	assert.Contains(t, result, "<timeout>30</timeout>")
}

// Test that specifically targets the strict mode interface{} conversions
func TestInterfaceTypesWithStrictMode(t *testing.T) {
	getDataWithInterfaces := func() interface{} {
		return map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{
					"nested": map[string]interface{}{
						"value": "deep",
					},
				},
			},
		}
	}
	
	opts := llml.Options{
		Strict: true,
		Prefix: "test",
	}
	
	var data interface{} = getDataWithInterfaces()
	result := llml.Sprintf(data, opts)
	
	// In strict mode, should use prefixes
	assert.Contains(t, result, "test-")
	assert.Contains(t, result, "<test-items>")
}