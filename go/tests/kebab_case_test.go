package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

func TestKebabCaseConversion(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user_name": "Alice",
		"userAge":   30,
	})
	expected := "<userAge>30</userAge>\n<user_name>Alice</user_name>"
	assert.Equal(t, expected, result)
}

func TestStringWithSpacesInKey(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"key with spaces": "value",
	})
	expected := "<key with spaces>value</key with spaces>"
	assert.Equal(t, expected, result)
}

// Advanced CamelCase Tests
func TestAdvancedCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"userName":  "Alice",
		"firstName": "Bob",
	})
	assert.Contains(t, result, "<userName>Alice</userName>")
	assert.Contains(t, result, "<firstName>Bob</firstName>")
}

func TestMultipleWordCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"getUserName": "function",
		"setUserAge":  "method",
	})
	assert.Contains(t, result, "<getUserName>function</getUserName>")
	assert.Contains(t, result, "<setUserAge>method</setUserAge>")
}

func TestAcronymCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"XMLHttpRequest": "api",
		"HTMLElement":    "dom",
	})
	assert.Contains(t, result, "<XMLHttpRequest>api</XMLHttpRequest>")
	assert.Contains(t, result, "<HTMLElement>dom</HTMLElement>")
}

func TestMixedCasesWithAcronyms(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"XMLParser":       "tool",
		"HTTPSConnection": "secure",
	})
	assert.Contains(t, result, "<XMLParser>tool</XMLParser>")
	assert.Contains(t, result, "<HTTPSConnection>secure</HTTPSConnection>")
}

func TestNumbersInCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user2Name":    "test",
		"config3Value": "data",
	})
	assert.Contains(t, result, "<user2Name>test</user2Name>")
	assert.Contains(t, result, "<config3Value>data</config3Value>")
}

func TestSingleLetterPrefixes(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"iPhone": "device",
		"iPad":   "tablet",
	})
	assert.Contains(t, result, "<iPhone>device</iPhone>")
	assert.Contains(t, result, "<iPad>tablet</iPad>")
}

func TestPreserveKebabCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user-name":  "Alice",
		"first-name": "Bob",
	})
	assert.Contains(t, result, "<user-name>Alice</user-name>")
	assert.Contains(t, result, "<first-name>Bob</first-name>")
}

func TestShortUppercaseSequences(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"A":   "single",
		"AB":  "double",
		"ABC": "triple",
	})
	assert.Contains(t, result, "<A>single</A>")
	assert.Contains(t, result, "<AB>double</AB>")
	assert.Contains(t, result, "<ABC>triple</ABC>")
}

func TestMixedPatterns(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"camelCase":   "test1",
		"snake_case":  "test2",
		"kebab-case":  "test3",
		"PascalCase":  "test4",
		"UPPER_SNAKE": "test5",
	})
	assert.Contains(t, result, "<camelCase>test1</camelCase>")
	assert.Contains(t, result, "<snake_case>test2</snake_case>")
	assert.Contains(t, result, "<kebab-case>test3</kebab-case>")
	assert.Contains(t, result, "<PascalCase>test4</PascalCase>")
	assert.Contains(t, result, "<UPPER_SNAKE>test5</UPPER_SNAKE>")
}
