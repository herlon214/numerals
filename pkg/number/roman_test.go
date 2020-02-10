package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanSystemName(t *testing.T) {
	numberSystem := NewRomanSystem()

	assert.Equal(t, numberSystem.Name(), "Roman")
}

func TestEncodeRoman(t *testing.T) {
	numberSystem := NewRomanSystem()

	var input int64 = 2020
	expected := "MMXX"

	result, success := numberSystem.Encode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func BenchmarkEncodeRoman(b *testing.B) {
	var input int64 = 2020

	for n := 0; n < b.N; n++ {
		numberSystem := NewRomanSystem()

		numberSystem.Encode(input)
	}
}

func TestDecodeRoman(t *testing.T) {
	numberSystem := NewRomanSystem()

	input := "MMXX"
	var expected int64 = 2020

	result, success := numberSystem.Decode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)

	input = "IV"
	expected = 4

	result, success = numberSystem.Decode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func BenchmarkDecodeRoman(b *testing.B) {
	input := "MMXX"

	for n := 0; n < b.N; n++ {
		numberSystem := NewRomanSystem()

		numberSystem.Decode(input)
	}
}

func TestFailDecodeRoman(t *testing.T) {
	numberSystem := NewRomanSystem()

	input := "123"
	var expected int64 = 0

	result, success := numberSystem.Decode(input)
	assert.False(t, success)
	assert.Equal(t, expected, result)
}

func TestFailEncodeRoman(t *testing.T) {
	numberSystem := NewRomanSystem()

	var input int64 = 0
	expected := ""

	result, success := numberSystem.Encode(input)
	assert.False(t, success)
	assert.Equal(t, expected, result)
}
