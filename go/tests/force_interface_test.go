package llml_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Use reflection to absolutely force interface{} types

func TestForceMapStringInterfaceWithReflection(t *testing.T) {
	// Create actual map[string]interface{} using reflection
	mapType := reflect.TypeOf((*map[string]interface{})(nil)).Elem()
	mapValue := reflect.MakeMap(mapType)
	
	// Add values using reflection
	mapValue.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf("value1"))
	mapValue.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(42))
	
	// Get interface{} that is actually map[string]interface{}
	data := mapValue.Interface()
	
	result := llml.Sprintf(data)
	expected := "<key1>value1</key1>\n<key2>42</key2>"
	assert.Equal(t, expected, result)
}

func TestForceSliceInterfaceWithReflection(t *testing.T) {
	// Create actual []interface{} using reflection
	sliceType := reflect.TypeOf((*[]interface{})(nil)).Elem()
	sliceValue := reflect.MakeSlice(sliceType, 0, 3)
	
	// Add values using reflection
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("item1"))
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("item2"))
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf(123))
	
	// Get interface{} that is actually []interface{}
	data := sliceValue.Interface()
	
	result := llml.Sprintf(data)
	expected := "<1>item1</1>\n<2>item2</2>\n<3>123</3>"
	assert.Equal(t, expected, result)
}

func TestForceNestedMapStringInterfaceInKeyValue(t *testing.T) {
	// Create nested map[string]interface{} using reflection
	nestedMapType := reflect.TypeOf((*map[string]interface{})(nil)).Elem()
	nestedMapValue := reflect.MakeMap(nestedMapType)
	nestedMapValue.SetMapIndex(reflect.ValueOf("nested"), reflect.ValueOf("value"))
	nestedMapValue.SetMapIndex(reflect.ValueOf("count"), reflect.ValueOf(5))
	
	// Put it in a map[string]any
	data := map[string]any{
		"config": nestedMapValue.Interface(),
	}
	
	result := llml.Sprintf(data)
	expected := "<config>\n  <count>5</count>\n  <nested>value</nested>\n</config>"
	assert.Equal(t, expected, result)
}

func TestForceSliceInterfaceInKeyValue(t *testing.T) {
	// Create []interface{} using reflection
	sliceType := reflect.TypeOf((*[]interface{})(nil)).Elem()
	sliceValue := reflect.MakeSlice(sliceType, 0, 3)
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("a"))
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("b"))
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("c"))
	
	// Put it in a map[string]any
	data := map[string]any{
		"items": sliceValue.Interface(),
	}
	
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n  <items-3>c</items-3>\n</items>"
	assert.Equal(t, expected, result)
}

func TestForceMapStringInterfaceInArray(t *testing.T) {
	// Create map[string]interface{} using reflection
	mapType := reflect.TypeOf((*map[string]interface{})(nil)).Elem()
	mapValue1 := reflect.MakeMap(mapType)
	mapValue1.SetMapIndex(reflect.ValueOf("id"), reflect.ValueOf(1))
	mapValue1.SetMapIndex(reflect.ValueOf("name"), reflect.ValueOf("test1"))
	
	mapValue2 := reflect.MakeMap(mapType)
	mapValue2.SetMapIndex(reflect.ValueOf("id"), reflect.ValueOf(2))
	mapValue2.SetMapIndex(reflect.ValueOf("name"), reflect.ValueOf("test2"))
	
	// Put them in a named array
	data := map[string]any{
		"items": []any{
			mapValue1.Interface(),
			mapValue2.Interface(),
		},
	}
	
	result := llml.Sprintf(data)
	expected := "<items>\n  <items-1>\n    <id>1</id>\n    <name>test1</name>\n  </items-1>\n  <items-2>\n    <id>2</id>\n    <name>test2</name>\n  </items-2>\n</items>"
	assert.Equal(t, expected, result)
}

func TestForceMapStringInterfaceInDirectArray(t *testing.T) {
	// Create map[string]interface{} using reflection
	mapType := reflect.TypeOf((*map[string]interface{})(nil)).Elem()
	mapValue1 := reflect.MakeMap(mapType)
	mapValue1.SetMapIndex(reflect.ValueOf("type"), reflect.ValueOf("user"))
	mapValue1.SetMapIndex(reflect.ValueOf("id"), reflect.ValueOf(1))
	
	mapValue2 := reflect.MakeMap(mapType)
	mapValue2.SetMapIndex(reflect.ValueOf("type"), reflect.ValueOf("admin"))
	mapValue2.SetMapIndex(reflect.ValueOf("id"), reflect.ValueOf(2))
	
	// Put them in a direct array
	data := []any{
		mapValue1.Interface(),
		mapValue2.Interface(),
	}
	
	result := llml.Sprintf(data)
	expected := "<1>\n  <1-id>1</1-id>\n  <1-type>user</1-type>\n</1>\n<2>\n  <2-id>2</2-id>\n  <2-type>admin</2-type>\n</2>"
	assert.Equal(t, expected, result)
}

func TestForceSliceInterfaceInDirectArray(t *testing.T) {
	// Create []interface{} using reflection
	sliceType := reflect.TypeOf((*[]interface{})(nil)).Elem()
	
	sliceValue1 := reflect.MakeSlice(sliceType, 0, 2)
	sliceValue1 = reflect.Append(sliceValue1, reflect.ValueOf("first"))
	sliceValue1 = reflect.Append(sliceValue1, reflect.ValueOf("array"))
	
	sliceValue2 := reflect.MakeSlice(sliceType, 0, 2)
	sliceValue2 = reflect.Append(sliceValue2, reflect.ValueOf("second"))
	sliceValue2 = reflect.Append(sliceValue2, reflect.ValueOf("array"))
	
	// Put them in a direct array
	data := []any{
		sliceValue1.Interface(),
		sliceValue2.Interface(),
	}
	
	result := llml.Sprintf(data)
	expected := "<1>\n  <1>first</1>\n  <2>array</2>\n</1>\n<2>\n  <1>second</1>\n  <2>array</2>\n</2>"
	assert.Equal(t, expected, result)
}