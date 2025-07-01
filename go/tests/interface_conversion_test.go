package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Test interface conversion paths for map[string]interface{} and []interface{}

func TestMapStringInterfaceConversion(t *testing.T) {
	// Test map[string]interface{} -> map[string]any conversion
	data := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}
	result := llml.Sprintf(data)
	expected := "<age>30</age>\n<name>Alice</name>"
	assert.Equal(t, expected, result)
}

func TestSliceInterfaceConversion(t *testing.T) {
	// Test []interface{} -> []any conversion for direct arrays
	data := []interface{}{"hello", 42, true}
	result := llml.Sprintf(data)
	expected := "<1>hello</1>\n<2>42</2>\n<3>true</3>"
	assert.Equal(t, expected, result)
}

func TestNestedMapStringInterfaceConversion(t *testing.T) {
	// Test nested map[string]interface{} conversion
	data := map[string]any{
		"user": map[string]interface{}{
			"name": "Bob",
			"age":  25,
		},
	}
	result := llml.Sprintf(data)
	expected := "<user>\n  <age>25</age>\n  <name>Bob</name>\n</user>"
	assert.Equal(t, expected, result)
}

func TestSliceInterfaceInKeyValue(t *testing.T) {
	// Test []interface{} in key-value pair (formatKeyValue path)
	data := map[string]any{
		"items": []interface{}{"apple", "banana", "cherry"},
	}
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>apple</items-1>\n  <items-2>banana</items-2>\n  <items-3>cherry</items-3>\n</items>"
	assert.Equal(t, expected, result)
}

func TestMapStringInterfaceInList(t *testing.T) {
	// Test map[string]interface{} objects within arrays (formatList path)
	data := map[string]any{
		"users": []any{
			map[string]interface{}{
				"name": "Alice",
				"age":  30,
			},
			map[string]interface{}{
				"name": "Bob", 
				"age":  25,
			},
		},
	}
	result := llml.Sprintf(data)
	expected := "<users>\n  <users-1>\n    <age>30</age>\n    <name>Alice</name>\n  </users-1>\n  <users-2>\n    <age>25</age>\n    <name>Bob</name>\n  </users-2>\n</users>"
	assert.Equal(t, expected, result)
}

func TestMapStringInterfaceInDirectArray(t *testing.T) {
	// Test map[string]interface{} in direct array (formatSlice path)
	data := []any{
		map[string]interface{}{
			"id":   1,
			"name": "Item 1",
		},
		map[string]interface{}{
			"id":   2,
			"name": "Item 2",
		},
	}
	result := llml.Sprintf(data)
	expected := "<1>\n  <1-id>1</1-id>\n  <1-name>Item 1</1-name>\n</1>\n<2>\n  <2-id>2</2-id>\n  <2-name>Item 2</2-name>\n</2>"
	assert.Equal(t, expected, result)
}

func TestSliceInterfaceInDirectArray(t *testing.T) {
	// Test []interface{} nested arrays in direct array (formatSlice path)
	data := []any{
		[]interface{}{"a", "b"},
		[]interface{}{"c", "d"},
	}
	result := llml.Sprintf(data)
	expected := "<1>\n  <1>a</1>\n  <2>b</2>\n</1>\n<2>\n  <1>c</1>\n  <2>d</2>\n</2>"
	assert.Equal(t, expected, result)
}

func TestEmptyMapStringInterface(t *testing.T) {
	// Test empty map[string]interface{} conversion
	data := map[string]interface{}{}
	result := llml.Sprintf(data)
	expected := ""
	assert.Equal(t, expected, result)
}

func TestEmptySliceInterface(t *testing.T) {
	// Test empty []interface{} conversion
	data := []interface{}{}
	result := llml.Sprintf(data)
	expected := ""
	assert.Equal(t, expected, result)
}

func TestEmptyMapStringInterfaceInDirectArray(t *testing.T) {
	// Test empty map[string]interface{} objects in direct arrays
	data := []any{
		map[string]interface{}{},
		map[string]interface{}{"name": "test"},
	}
	result := llml.Sprintf(data)
	expected := "<1></1>\n<2>\n  <2-name>test</2-name>\n</2>"
	assert.Equal(t, expected, result)
}