package number

import (
	"strconv"
)

type binary struct{}

// NewBinarySystem creates the the numeric system that represents
// the binary system
func NewBinarySystem() System {
	return &binary{}
}

func (b *binary) Name() string {
	return "Binary"
}

func (b *binary) Decode(input string) (int64, bool) {
	number, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		return 0, false
	}

	return number, true
}

func (b *binary) Encode(input int64) (string, bool) {
	return strconv.FormatInt(input, 2), true
}
