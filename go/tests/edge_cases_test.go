package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Test edge cases and error conditions

func TestToKebabCaseEmptyString(t *testing.T) {
	// Test empty string edge case in toKebabCase (this is internal, so test via key)
	result := llml.Sprintf(map[string]any{
		"": "value",
	})
	expected := "<>value</>"
	assert.Equal(t, expected, result)
}

func TestDeeplyNestedStructures(t *testing.T) {
	// Test deeply nested structures (3+ levels)
	data := map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": map[string]any{
					"level4": "deep value",
				},
			},
		},
	}
	result := llml.Sprintf(data)
	expected := "<level1>  <level2>    <level3>      <level4>deep value</level4></level3></level2></level1>"
	assert.Equal(t, expected, result)
}

func TestEmptyArraysInDifferentContexts(t *testing.T) {
	// Test empty arrays in various contexts to ensure they're properly skipped
	data := map[string]any{
		"emptyItems": []any{},
		"normalItems": []any{"item1", "item2"},
		"mixedItems": []any{
			map[string]any{"name": "test"},
			[]any{}, // Empty nested array (creates empty tag)
		},
	}
	result := llml.Sprintf(data)
	expected := "<mixedItems>\n  <mixedItems-1>\n    <name>test</name>\n  </mixedItems-1>\n  <mixedItems-2></mixedItems-2>\n</mixedItems>\n<normalItems>\n  <normalItems-1>item1</normalItems-1>\n  <normalItems-2>item2</normalItems-2>\n</normalItems>"
	assert.Equal(t, expected, result)
}

func TestComplexArrayScenariosInDirectArray(t *testing.T) {
	// Test complex scenarios for formatSlice function coverage
	data := []any{
		"simple string",
		map[string]any{
			"nested": "object",
		},
		[]any{"nested", "array"},
		[]any{}, // Empty array (should be skipped)
		map[string]any{}, // Empty object
	}
	result := llml.Sprintf(data)
	expected := "<1>simple string</1>\n<2>\n  <2-nested>object</2-nested>\n</2>\n<3>\n  <1>nested</1>\n  <2>array</2>\n</3>\n<5></5>"
	assert.Equal(t, expected, result)
}

func TestMultilineStringWithDifferentFormats(t *testing.T) {
	// Test multiline strings in various contexts
	multilineText := "Line 1\nLine 2\nLine 3"
	
	// In simple key-value
	result1 := llml.Sprintf(map[string]any{
		"text": multilineText,
	})
	expected1 := "<text>\n  Line 1\n  Line 2\n  Line 3\n</text>"
	assert.Equal(t, expected1, result1)
	
	// In nested context - Go implementation doesn't add extra indentation to multiline content
	result2 := llml.Sprintf(map[string]any{
		"container": map[string]any{
			"text": multilineText,
		},
	})
	expected2 := "<container>\n  <text>\n  Line 1\n  Line 2\n  Line 3\n  </text>\n</container>"
	assert.Equal(t, expected2, result2)
}

func TestSpecialCharactersInKeys(t *testing.T) {
	// Test special characters and edge cases in key names
	data := map[string]any{
		"key with spaces": "value1",
		"key_with_underscores": "value2", 
		"keyWithCamelCase": "value3",
		"KEY_WITH_CAPS": "value4",
		"key-already-kebab": "value5",
		"XMLHttpRequest": "value6", // Test acronym handling
	}
	result := llml.Sprintf(data)
	// Keys are sorted alphabetically, so adjust expected order
	expected := "<KEY_WITH_CAPS>value4</KEY_WITH_CAPS>\n<XMLHttpRequest>value6</XMLHttpRequest>\n<key with spaces>value1</key with spaces>\n<key-already-kebab>value5</key-already-kebab>\n<keyWithCamelCase>value3</keyWithCamelCase>\n<key_with_underscores>value2</key_with_underscores>"
	assert.Equal(t, expected, result)
}

func TestZeroValuesForAllTypes(t *testing.T) {
	// Test zero values for all types to ensure they're properly handled
	data := map[string]any{
		"emptyString": "",
		"zeroInt": 0,
		"zeroFloat": 0.0,
		"falseBool": false,
		"nilValue": nil,
	}
	result := llml.Sprintf(data)
	expected := "<emptyString></emptyString>\n<falseBool>false</falseBool>\n<nilValue>nil</nilValue>\n<zeroFloat>0</zeroFloat>\n<zeroInt>0</zeroInt>"
	assert.Equal(t, expected, result)
}

func TestLargeDataStructures(t *testing.T) {
	// Test with larger data structures to ensure performance
	items := make([]any, 50)
	for i := 0; i < 50; i++ {
		items[i] = map[string]any{
			"id": i,
			"name": "Item " + string(rune(i+'A')),
		}
	}
	
	data := map[string]any{
		"items": items,
	}
	
	result := llml.Sprintf(data)
	// Just check that it doesn't crash and produces reasonable output
	assert.Contains(t, result, "<items>")
	assert.Contains(t, result, "<items-1>")
	assert.Contains(t, result, "<items-50>")
	assert.Contains(t, result, "</items>")
}