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
	expected := "<user_tasks>\n" +
		"  <user_tasks-1>task1</user_tasks-1>\n" +
		"  <user_tasks-2>task2</user_tasks-2>\n" +
		"</user_tasks>"
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
	assert.Contains(t, result, "<userTasks>")
	assert.Contains(t, result, "<userTasks-1>task1</userTasks-1>")
	assert.Contains(t, result, "<userTasks-2>task2</userTasks-2>")
	assert.Contains(t, result, "</userTasks>")
	assert.Contains(t, result, "<XMLElements>")
	assert.Contains(t, result, "<XMLElements-1>element1</XMLElements-1>")
	assert.Contains(t, result, "<XMLElements-2>element2</XMLElements-2>")
	assert.Contains(t, result, "</XMLElements>")
}
