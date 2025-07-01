package llml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenbase-ai/llml/go/pkg/llml"
)

// Test comprehensive type coverage for all primitive types

func TestInt8Type(t *testing.T) {
	var val int8 = 127
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>127</value>"
	assert.Equal(t, expected, result)
}

func TestInt16Type(t *testing.T) {
	var val int16 = 32767
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>32767</value>"
	assert.Equal(t, expected, result)
}

func TestInt32Type(t *testing.T) {
	var val int32 = 2147483647
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>2147483647</value>"
	assert.Equal(t, expected, result)
}

func TestInt64Type(t *testing.T) {
	var val int64 = 9223372036854775807
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>9223372036854775807</value>"
	assert.Equal(t, expected, result)
}

func TestUintType(t *testing.T) {
	var val uint = 4294967295
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>4294967295</value>"
	assert.Equal(t, expected, result)
}

func TestUint8Type(t *testing.T) {
	var val uint8 = 255
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>255</value>"
	assert.Equal(t, expected, result)
}

func TestUint16Type(t *testing.T) {
	var val uint16 = 65535
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>65535</value>"
	assert.Equal(t, expected, result)
}

func TestUint32Type(t *testing.T) {
	var val uint32 = 4294967295
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>4294967295</value>"
	assert.Equal(t, expected, result)
}

func TestUint64Type(t *testing.T) {
	var val uint64 = 18446744073709551615
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>18446744073709551615</value>"
	assert.Equal(t, expected, result)
}

func TestFloat32Type(t *testing.T) {
	var val float32 = 3.14159
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>3.14159</value>"
	assert.Equal(t, expected, result)
}

func TestUnknownType(t *testing.T) {
	type CustomStruct struct {
		Name string
	}
	val := CustomStruct{Name: "test"}
	result := llml.Sprintf(map[string]any{
		"value": val,
	})
	expected := "<value>{test}</value>"
	assert.Equal(t, expected, result)
}

func TestDirectPrimitiveTypes(t *testing.T) {
	// Test direct primitive calls without wrapping in map
	
	var int8Val int8 = 127
	assert.Equal(t, "127", llml.Sprintf(int8Val))
	
	var int16Val int16 = 32767
	assert.Equal(t, "32767", llml.Sprintf(int16Val))
	
	var int32Val int32 = 2147483647
	assert.Equal(t, "2147483647", llml.Sprintf(int32Val))
	
	var int64Val int64 = 9223372036854775807
	assert.Equal(t, "9223372036854775807", llml.Sprintf(int64Val))
	
	var uintVal uint = 4294967295
	assert.Equal(t, "4294967295", llml.Sprintf(uintVal))
	
	var uint8Val uint8 = 255
	assert.Equal(t, "255", llml.Sprintf(uint8Val))
	
	var uint16Val uint16 = 65535
	assert.Equal(t, "65535", llml.Sprintf(uint16Val))
	
	var uint32Val uint32 = 4294967295
	assert.Equal(t, "4294967295", llml.Sprintf(uint32Val))
	
	var uint64Val uint64 = 18446744073709551615
	assert.Equal(t, "18446744073709551615", llml.Sprintf(uint64Val))
	
	var float32Val float32 = 3.14159
	assert.Equal(t, "3.14159", llml.Sprintf(float32Val))
}