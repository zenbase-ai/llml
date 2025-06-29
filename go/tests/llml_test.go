package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml-go/pkg/llml"
)

func TestNoArgs(t *testing.T) {
	result := llml.LLML(nil)
	expected := "nil"
	assert.Equal(t, expected, result)
}

func TestEmptyMap(t *testing.T) {
	result := llml.LLML(map[string]any{})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestEmptySlice(t *testing.T) {
	result := llml.LLML([]any{})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestEmptyString(t *testing.T) {
	result := llml.LLML(map[string]any{
		"empty": "",
	})
	expected := "<empty></empty>"
	assert.Equal(t, expected, result)
}

func TestZeroValue(t *testing.T) {
	result := llml.LLML(map[string]any{
		"zero": 0,
	})
	expected := "<zero>0</zero>"
	assert.Equal(t, expected, result)
}

func TestFalseBoolean(t *testing.T) {
	result := llml.LLML(map[string]any{
		"disabled": false,
	})
	expected := "<disabled>false</disabled>"
	assert.Equal(t, expected, result)
}

func TestNilValue(t *testing.T) {
	result := llml.LLML(map[string]any{
		"nothing": nil,
	})
	expected := "<nothing>nil</nothing>"
	assert.Equal(t, expected, result)
}

func TestSimpleStringValue(t *testing.T) {
	result := llml.LLML(map[string]any{
		"instructions": "Follow these steps",
	})
	expected := "<instructions>Follow these steps</instructions>"
	assert.Equal(t, expected, result)
}

func TestIntegerValue(t *testing.T) {
	result := llml.LLML(map[string]any{
		"count": 42,
	})
	expected := "<count>42</count>"
	assert.Equal(t, expected, result)
}

func TestFloatValue(t *testing.T) {
	result := llml.LLML(map[string]any{
		"temperature": 98.6,
	})
	expected := "<temperature>98.6</temperature>"
	assert.Equal(t, expected, result)
}

func TestBooleanValue(t *testing.T) {
	result := llml.LLML(map[string]any{
		"enabled": true,
	})
	expected := "<enabled>true</enabled>"
	assert.Equal(t, expected, result)
}

func TestKebabCaseConversion(t *testing.T) {
	result := llml.LLML(map[string]any{
		"user_name": "Alice",
		"userAge":   30,
	})
	expected := "<user-age>30</user-age>\n<user-name>Alice</user-name>"
	assert.Equal(t, expected, result)
}

func TestStringWithSpacesInKey(t *testing.T) {
	result := llml.LLML(map[string]any{
		"key with spaces": "value",
	})
	expected := "<key-with-spaces>value</key-with-spaces>"
	assert.Equal(t, expected, result)
}

func TestMultipleSimpleValues(t *testing.T) {
	result := llml.LLML(map[string]any{
		"name":   "Alice",
		"age":    30,
		"active": true,
	})
	// Note: Map iteration order is not guaranteed in Go, so we need to check for either order
	// For this test, we'll assume a specific order or check that all parts are present
	assert.Contains(t, result, "<name>Alice</name>")
	assert.Contains(t, result, "<age>30</age>")
	assert.Contains(t, result, "<active>true</active>")
}

func TestBasicIndentation(t *testing.T) {
	result := llml.LLML(map[string]any{
		"message": "Hello",
	}, llml.Options{Indent: "  "})
	expected := "  <message>Hello</message>"
	assert.Equal(t, expected, result)
}

func TestWithPrefix(t *testing.T) {
	result := llml.LLML(map[string]any{
		"config": "value",
	}, llml.Options{Prefix: "app"})
	expected := "<app-config>value</app-config>"
	assert.Equal(t, expected, result)
}

func TestMultilineContent(t *testing.T) {
	content := `
    Line 1
    Line 2
    Line 3
    `
	result := llml.LLML(map[string]any{
		"description": content,
	})
	expected := "<description>\n  Line 1\n  Line 2\n  Line 3\n</description>"
	assert.Equal(t, expected, result)
}

func TestEmptyList(t *testing.T) {
	result := llml.LLML(map[string]any{
		"items": []any{},
	})
	expected := "<items-list></items-list>"
	assert.Equal(t, expected, result)
}

func TestSimpleListWithWrapper(t *testing.T) {
	result := llml.LLML(map[string]any{
		"rules": []any{"first", "second", "third"},
	})
	expected := "<rules-list>\n" +
		"  <rules-1>first</rules-1>\n" +
		"  <rules-2>second</rules-2>\n" +
		"  <rules-3>third</rules-3>\n" +
		"</rules-list>"
	assert.Equal(t, expected, result)
}

func TestListWithNumbers(t *testing.T) {
	result := llml.LLML(map[string]any{
		"numbers": []any{1, 2, 3},
	})
	expected := "<numbers-list>\n" +
		"  <numbers-1>1</numbers-1>\n" +
		"  <numbers-2>2</numbers-2>\n" +
		"  <numbers-3>3</numbers-3>\n" +
		"</numbers-list>"
	assert.Equal(t, expected, result)
}

func TestListKebabCaseConversion(t *testing.T) {
	result := llml.LLML(map[string]any{
		"user_tasks": []any{"task1", "task2"},
	})
	expected := "<user-tasks-list>\n" +
		"  <user-tasks-1>task1</user-tasks-1>\n" +
		"  <user-tasks-2>task2</user-tasks-2>\n" +
		"</user-tasks-list>"
	assert.Equal(t, expected, result)
}

func TestListWithIndentation(t *testing.T) {
	result := llml.LLML(map[string]any{
		"items": []any{"a", "b"},
	}, llml.Options{Indent: "  "})
	expected := "  <items-list>\n" +
		"    <items-1>a</items-1>\n" +
		"    <items-2>b</items-2>\n" +
		"  </items-list>"
	assert.Equal(t, expected, result)
}

func TestListWithPrefix(t *testing.T) {
	result := llml.LLML(map[string]any{
		"items": []any{"a", "b"},
	}, llml.Options{Prefix: "app"})
	expected := "<app-items-list>\n" +
		"  <app-items-1>a</app-items-1>\n" +
		"  <app-items-2>b</app-items-2>\n" +
		"</app-items-list>"
	assert.Equal(t, expected, result)
}

func TestNestedDict(t *testing.T) {
	result := llml.LLML(map[string]any{
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
	result := llml.LLML(map[string]any{
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

func TestListOfDicts(t *testing.T) {
	result := llml.LLML(map[string]any{
		"data": []any{
			map[string]any{"name": "Alice", "age": 30},
			map[string]any{"name": "Bob", "age": 25},
		},
	})
	assert.Contains(t, result, "<data-list>")
	assert.Contains(t, result, "</data-list>")
	assert.Contains(t, result, "<data-1>")
	assert.Contains(t, result, "</data-1>")
	assert.Contains(t, result, "<data-2>")
	assert.Contains(t, result, "</data-2>")
	assert.Contains(t, result, "<data-1-name>Alice</data-1-name>")
	assert.Contains(t, result, "<data-1-age>30</data-1-age>")
	assert.Contains(t, result, "<data-2-name>Bob</data-2-name>")
	assert.Contains(t, result, "<data-2-age>25</data-2-age>")
}

func TestMixedContent(t *testing.T) {
	result := llml.LLML(map[string]any{
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
