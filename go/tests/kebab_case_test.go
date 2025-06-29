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
	expected := "<user-age>30</user-age>\n<user-name>Alice</user-name>"
	assert.Equal(t, expected, result)
}

func TestStringWithSpacesInKey(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"key with spaces": "value",
	})
	expected := "<key-with-spaces>value</key-with-spaces>"
	assert.Equal(t, expected, result)
}

// Advanced CamelCase Tests
func TestAdvancedCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"userName":  "Alice",
		"firstName": "Bob",
	})
	assert.Contains(t, result, "<user-name>Alice</user-name>")
	assert.Contains(t, result, "<first-name>Bob</first-name>")
}

func TestMultipleWordCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"getUserName": "function",
		"setUserAge":  "method",
	})
	assert.Contains(t, result, "<get-user-name>function</get-user-name>")
	assert.Contains(t, result, "<set-user-age>method</set-user-age>")
}

func TestAcronymCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"XMLHttpRequest": "api",
		"HTMLElement":    "dom",
	})
	assert.Contains(t, result, "<xml-http-request>api</xml-http-request>")
	assert.Contains(t, result, "<html-element>dom</html-element>")
}

func TestMixedCasesWithAcronyms(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"XMLParser":       "tool",
		"HTTPSConnection": "secure",
	})
	assert.Contains(t, result, "<xml-parser>tool</xml-parser>")
	assert.Contains(t, result, "<https-connection>secure</https-connection>")
}

func TestNumbersInCamelCase(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"user2Name":    "test",
		"config3Value": "data",
	})
	assert.Contains(t, result, "<user2-name>test</user2-name>")
	assert.Contains(t, result, "<config3-value>data</config3-value>")
}

func TestSingleLetterPrefixes(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"iPhone": "device",
		"iPad":   "tablet",
	})
	assert.Contains(t, result, "<i-phone>device</i-phone>")
	assert.Contains(t, result, "<i-pad>tablet</i-pad>")
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
	assert.Contains(t, result, "<a>single</a>")
	assert.Contains(t, result, "<ab>double</ab>")
	assert.Contains(t, result, "<abc>triple</abc>")
}

func TestMixedPatterns(t *testing.T) {
	result := llml.Sprintf(map[string]any{
		"camelCase":   "test1",
		"snake_case":  "test2",
		"kebab-case":  "test3",
		"PascalCase":  "test4",
		"UPPER_SNAKE": "test5",
	})
	assert.Contains(t, result, "<camel-case>test1</camel-case>")
	assert.Contains(t, result, "<snake-case>test2</snake-case>")
	assert.Contains(t, result, "<kebab-case>test3</kebab-case>")
	assert.Contains(t, result, "<pascal-case>test4</pascal-case>")
	assert.Contains(t, result, "<upper-snake>test5</upper-snake>")
}
