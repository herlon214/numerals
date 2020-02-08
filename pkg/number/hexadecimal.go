package number

import (
	"fmt"
	"strconv"
)

type hexadecimal struct{}

// NewHexadecimalSystem creates the the numeric system that represents
// the decimal system
func NewHexadecimalSystem() System {
	return &hexadecimal{}
}

func (h *hexadecimal) Name() string {
	return "Hexadecimal"
}

func (h *hexadecimal) Decode(input string) (int64, bool) {
	number, err := strconv.ParseInt(input, 16, 64)
	if err != nil {
		return 0, false
	}

	return number, true
}

func (h *hexadecimal) Encode(input int64) (string, bool) {
	return fmt.Sprintf("%X", input), true
}
