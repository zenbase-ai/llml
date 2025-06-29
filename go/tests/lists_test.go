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
	expected := "<items-list></items-list>"
	assert.Equal(t, expected, result)
}

func TestSimpleListWithWrapper(t *testing.T) {
	result := llml.Sprintf(map[string]any{
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
	result := llml.Sprintf(map[string]any{
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
	result := llml.Sprintf(map[string]any{
		"user_tasks": []any{"task1", "task2"},
	})
	expected := "<user-tasks-list>\n" +
		"  <user-tasks-1>task1</user-tasks-1>\n" +
		"  <user-tasks-2>task2</user-tasks-2>\n" +
		"</user-tasks-list>"
	assert.Equal(t, expected, result)
}

func TestListOfDicts(t *testing.T) {
	result := llml.Sprintf(map[string]any{
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

func TestCamelCaseInArrayKeys(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"userTasks":   []any{"task1", "task2"},
		"XMLElements": []any{"element1", "element2"},
	})
	assert.Contains(t, result, "<user-tasks-list>")
	assert.Contains(t, result, "<user-tasks-1>task1</user-tasks-1>")
	assert.Contains(t, result, "<user-tasks-2>task2</user-tasks-2>")
	assert.Contains(t, result, "</user-tasks-list>")
	assert.Contains(t, result, "<xml-elements-list>")
	assert.Contains(t, result, "<xml-elements-1>element1</xml-elements-1>")
	assert.Contains(t, result, "<xml-elements-2>element2</xml-elements-2>")
	assert.Contains(t, result, "</xml-elements-list>")
}
