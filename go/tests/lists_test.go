package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestEmptyList(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"items": []any{},
	})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestSimpleListWithWrapper(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"rules": []any{"first", "second", "third"},
	})
	expected := "<rules>\n" +
		"  <rules-1>first</rules-1>\n" +
		"  <rules-2>second</rules-2>\n" +
		"  <rules-3>third</rules-3>\n" +
		"</rules>"
	assert.Equal(t, expected, result)
}

func TestListWithNumbers(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"numbers": []any{1, 2, 3},
	})
	expected := "<numbers>\n" +
		"  <numbers-1>1</numbers-1>\n" +
		"  <numbers-2>2</numbers-2>\n" +
		"  <numbers-3>3</numbers-3>\n" +
		"</numbers>"
	assert.Equal(t, expected, result)
}

func TestListKebabCaseConversion(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user_tasks": []any{"task1", "task2"},
	})
	expected := "<user-tasks>\n" +
		"  <user-tasks-1>task1</user-tasks-1>\n" +
		"  <user-tasks-2>task2</user-tasks-2>\n" +
		"</user-tasks>"
	assert.Equal(t, expected, result)
}

func TestListOfDicts(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"data": []any{
			map[string]any{"name": "Alice", "age": 30},
			map[string]any{"name": "Bob", "age": 25},
		},
	})
	assert.Contains(t, result, "<data>")
	assert.Contains(t, result, "</data>")
	assert.Contains(t, result, "<data-1>")
	assert.Contains(t, result, "</data-1>")
	assert.Contains(t, result, "<data-2>")
	assert.Contains(t, result, "</data-2>")
	// In non-strict mode (default), nested object properties don't include parent prefixes
	assert.Contains(t, result, "<name>Alice</name>")
	assert.Contains(t, result, "<age>30</age>")
	assert.Contains(t, result, "<name>Bob</name>")
	assert.Contains(t, result, "<age>25</age>")
}

func TestCamelCaseInArrayKeys(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"userTasks":   []any{"task1", "task2"},
		"XMLElements": []any{"element1", "element2"},
	})
	assert.Contains(t, result, "<user-tasks>")
	assert.Contains(t, result, "<user-tasks-1>task1</user-tasks-1>")
	assert.Contains(t, result, "<user-tasks-2>task2</user-tasks-2>")
	assert.Contains(t, result, "</user-tasks>")
	assert.Contains(t, result, "<xml-elements>")
	assert.Contains(t, result, "<xml-elements-1>element1</xml-elements-1>")
	assert.Contains(t, result, "<xml-elements-2>element2</xml-elements-2>")
	assert.Contains(t, result, "</xml-elements>")
}
