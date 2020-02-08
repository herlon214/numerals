package number

import "strconv"

type decimal struct{}

// NewDecimalSystem creates the the numeric system that represents
// the decimal system
func NewDecimalSystem() System {
	return &decimal{}
}

func (d *decimal) Name() string {
	return "Decimal"
}

func (d *decimal) Decode(input string) (int64, bool) {
	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, false
	}

	return number, true
}

func (d *decimal) Encode(input int64) (string, bool) {
	return strconv.FormatInt(input, 10), true
}
