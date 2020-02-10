package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySystemName(t *testing.T) {
	numberSystem := NewBinarySystem()

	assert.Equal(t, numberSystem.Name(), "Binary")
}

func TestEncodeBinary(t *testing.T) {
	numberSystem := NewBinarySystem()

	var input int64 = 123123123
	expected := "111010101101011010110110011"

	result, success := numberSystem.Encode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func BenchmarkEncodeBinary(b *testing.B) {
	var input int64 = 123123123

	for n := 0; n < b.N; n++ {
		numberSystem := NewBinarySystem()

		numberSystem.Encode(input)
	}
}

func TestDecodeBinary(t *testing.T) {
	numberSystem := NewBinarySystem()

	input := "111010101101011010110110011"
	var expected int64 = 123123123

	result, success := numberSystem.Decode(input)
	assert.True(t, success)
	assert.Equal(t, expected, result)
}

func BenchmarkDecodeBinary(b *testing.B) {
	input := "111010101101011010110110011"

	for n := 0; n < b.N; n++ {
		numberSystem := NewBinarySystem()

		numberSystem.Decode(input)
	}
}

func TestFailDecodeBinary(t *testing.T) {
	numberSystem := NewBinarySystem()

	input := "POIUYTR"
	var expected int64 = 0

	result, success := numberSystem.Decode(input)
	assert.False(t, success)
	assert.Equal(t, expected, result)
}
