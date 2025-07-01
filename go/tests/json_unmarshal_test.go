package llml_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Use JSON unmarshaling to force actual interface{} types at runtime

func TestJSONUnmarshalMapStringInterface(t *testing.T) {
	// JSON unmarshaling into map[string]interface{} forces the actual type
	jsonData := `{"name": "test", "value": 42}`
	
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	assert.NoError(t, err)
	
	// Now data is truly map[string]interface{}, not map[string]any
	result := llml.Sprintf(data)
	expected := "<name>test</name>\n<value>42</value>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalSliceInterface(t *testing.T) {
	// JSON unmarshaling into []interface{} forces the actual type
	jsonData := `["item1", "item2", 123]`
	
	var data []interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	assert.NoError(t, err)
	
	// Now data is truly []interface{}, not []any
	result := llml.Sprintf(data)
	expected := "<1>item1</1>\n<2>item2</2>\n<3>123</3>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalNestedMapStringInterfaceInKeyValue(t *testing.T) {
	// Create structure with map[string]interface{} as nested value
	jsonData := `{"config": {"nested": "value", "count": 5}}`
	
	var outerData map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &outerData)
	assert.NoError(t, err)
	
	// Put the nested map[string]interface{} into a map[string]any
	data := map[string]any{
		"config": outerData["config"], // This is map[string]interface{}
	}
	
	result := llml.Sprintf(data)
	expected := "<config>\n  <count>5</count>\n  <nested>value</nested>\n</config>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalSliceInterfaceInKeyValue(t *testing.T) {
	// Create structure with []interface{} as nested value
	jsonData := `{"items": ["a", "b", "c"]}`
	
	var outerData map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &outerData)
	assert.NoError(t, err)
	
	// Put the []interface{} into a map[string]any
	data := map[string]any{
		"items": outerData["items"], // This is []interface{}
	}
	
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n  <items-3>c</items-3>\n</items>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalMapStringInterfaceInArray(t *testing.T) {
	// Create array with map[string]interface{} objects
	jsonData := `[{"id": 1, "name": "test1"}, {"id": 2, "name": "test2"}]`
	
	var arrayData []interface{}
	err := json.Unmarshal([]byte(jsonData), &arrayData)
	assert.NoError(t, err)
	
	// Put the array with map[string]interface{} objects in a named array
	data := map[string]any{
		"items": arrayData, // Contains map[string]interface{} objects
	}
	
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>\n    <id>1</id>\n    <name>test1</name>\n  </items-1>\n  <items-2>\n    <id>2</id>\n    <name>test2</name>\n  </items-2>\n</items>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalMapStringInterfaceInDirectArray(t *testing.T) {
	// Create direct array with map[string]interface{} objects
	jsonData := `[{"type": "user", "id": 1}, {"type": "admin", "id": 2}]`
	
	var data []interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	assert.NoError(t, err)
	
	// data now contains actual map[string]interface{} objects
	result := llml.Sprintf(data)
	expected := "<1>\n  <1-id>1</1-id>\n  <1-type>user</1-type>\n</1>\n<2>\n  <2-id>2</2-id>\n  <2-type>admin</2-type>\n</2>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalSliceInterfaceInDirectArray(t *testing.T) {
	// Create direct array with []interface{} nested arrays
	jsonData := `[["first", "array"], ["second", "array"]]`
	
	var data []interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	assert.NoError(t, err)
	
	// data now contains actual []interface{} slices
	result := llml.Sprintf(data)
	expected := "<1>\n  <1>first</1>\n  <2>array</2>\n</1>\n<2>\n  <1>second</1>\n  <2>array</2>\n</2>"
	assert.Equal(t, expected, result)
}

func TestJSONUnmarshalComplexInterfaceStructure(t *testing.T) {
	// Complex JSON structure that forces multiple interface{} types
	jsonData := `{
		"users": [
			{
				"id": 1,
				"tags": ["admin", "active"],
				"config": {"timeout": 30}
			}
		],
		"settings": {
			"enabled": true,
			"items": ["a", "b"]
		}
	}`
	
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	assert.NoError(t, err)
	
	// This should force all the interface{} conversion paths
	result := llml.Sprintf(data)
	
	// Just verify it contains expected structure
	assert.Contains(t, result, "<users>")
	assert.Contains(t, result, "<settings>")
	assert.Contains(t, result, "<config>")
	assert.Contains(t, result, "<tags>")
}