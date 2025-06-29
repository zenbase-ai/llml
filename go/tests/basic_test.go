package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestNoArgs(t *testing.T) {
	result := llml.Sprintf(nil)
	expected := "nil"
	assert.Equal(t, expected, result)
}

func TestEmptyMap(t *testing.T) {
	result := llml.Sprintf(map[string]any{})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestEmptySlice(t *testing.T) {
	result := llml.Sprintf([]any{})
	expected := ""
	assert.Equal(t, expected, result)
}

func TestEmptyString(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"empty": "",
	})
	expected := "<empty></empty>"
	assert.Equal(t, expected, result)
}

func TestZeroValue(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"zero": 0,
	})
	expected := "<zero>0</zero>"
	assert.Equal(t, expected, result)
}

func TestFalseBoolean(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"disabled": false,
	})
	expected := "<disabled>false</disabled>"
	assert.Equal(t, expected, result)
}

func TestNilValue(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"nothing": nil,
	})
	expected := "<nothing>nil</nothing>"
	assert.Equal(t, expected, result)
}

func TestSimpleStringValue(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"instructions": "Follow these steps",
	})
	expected := "<instructions>Follow these steps</instructions>"
	assert.Equal(t, expected, result)
}

func TestIntegerValue(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"count": 42,
	})
	expected := "<count>42</count>"
	assert.Equal(t, expected, result)
}

func TestFloatValue(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"temperature": 98.6,
	})
	expected := "<temperature>98.6</temperature>"
	assert.Equal(t, expected, result)
}

func TestBooleanValue(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"enabled": true,
	})
	expected := "<enabled>true</enabled>"
	assert.Equal(t, expected, result)
}

func TestMultipleSimpleValues(t *testing.T) {
	result := llml.Sprintf(map[string]any{
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
