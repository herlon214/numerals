package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexadecimalSystemName(t *testing.T) {
	numberSystem := NewHexadecimalSystem()

	assert.Equal(t, numberSystem.Name(), "Hexadecimal")
}
func TestEncodeHexadecimal(t *testing.T) {
	numberSystem := NewHexadecimalSystem()

	var input int64 = 123123123
	expected := "756B5B3"

	result, success := numberSystem.Encode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func TestDecodeHexadecimal(t *testing.T) {
	numberSystem := NewHexadecimalSystem()

	input := "756B5B3"
	var expected int64 = 123123123

	result, success := numberSystem.Decode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func TestFailDecodeHexadecimal(t *testing.T) {
	numberSystem := NewHexadecimalSystem()

	input := "POIUYTR"
	var expected int64 = 0

	result, success := numberSystem.Decode(input)
	assert.False(t, success)
	assert.Equal(t, expected, result)
}
