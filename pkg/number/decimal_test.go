package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalSystemName(t *testing.T) {
	numberSystem := NewDecimalSystem()

	assert.Equal(t, numberSystem.Name(), "Decimal")
}

func TestEncodeDecimal(t *testing.T) {
	numberSystem := NewDecimalSystem()

	var input int64 = 123123123
	expected := "123123123"

	result, success := numberSystem.Encode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func TestDecodeDecimal(t *testing.T) {
	numberSystem := NewDecimalSystem()

	input := "123123123"
	var expected int64 = 123123123

	result, success := numberSystem.Decode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func TestFailDecodeDecimal(t *testing.T) {
	numberSystem := NewDecimalSystem()

	input := "abcde"
	var expected int64 = 0

	result, success := numberSystem.Decode(input)
	assert.False(t, success)
	assert.Equal(t, expected, result)
}
