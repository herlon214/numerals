package number

var romanValues = []int64{
	1000, 900, 500, 400,
	100, 90, 50, 40,
	10, 9, 5, 4, 1,
}

var romanSymbols = []string{
	"M", "CM", "D", "CD",
	"C", "XC", "L", "XL",
	"X", "IX", "V", "IV",
	"I",
}

var romanSymbolsDecimal = map[string]int64{
	romanSymbols[0]:  romanValues[0],
	romanSymbols[1]:  romanValues[1],
	romanSymbols[2]:  romanValues[2],
	romanSymbols[3]:  romanValues[3],
	romanSymbols[4]:  romanValues[4],
	romanSymbols[5]:  romanValues[5],
	romanSymbols[6]:  romanValues[6],
	romanSymbols[7]:  romanValues[7],
	romanSymbols[8]:  romanValues[8],
	romanSymbols[9]:  romanValues[9],
	romanSymbols[10]: romanValues[10],
	romanSymbols[11]: romanValues[11],
	romanSymbols[12]: romanValues[12],
}

type roman struct{}

// NewRomanSystem creates the the numeric system that represents
// the roman system
func NewRomanSystem() System {
	return &roman{}
}

func (r *roman) Name() string {
	return "Roman"
}

func (r *roman) Decode(input string) (int64, bool) {
	// Check if the input is valid
	if !r.isValid(input) {
		return 0, false
	}

	var total int64 = 0

	currentSymbol := ""
	nextSymbol := ""
	var currentValue int64 = 0
	var nextValue int64 = 0

	for i := 0; i < len(input); i++ {
		currentSymbol = string(input[i])
		currentValue = romanSymbolsDecimal[currentSymbol]

		// Check for subtractions first
		if i+1 < len(input) {
			nextSymbol = string(input[i+1])
			nextValue = romanSymbolsDecimal[nextSymbol]

			if currentValue < nextValue {
				total -= currentValue
			} else {
				total += currentValue
			}
		} else {
			total += currentValue
		}

	}

	return total, true
}

func (r *roman) Encode(input int64) (string, bool) {
	outputNumber := ""
	i := 0

	for input > 0 {
		k := int(input / romanValues[i])
		for j := 0; j < k; j++ {
			outputNumber += romanSymbols[i]
			input -= romanValues[i]
		}
		i++
	}

	return outputNumber, true
}

// Validate the input
func (r *roman) isValid(input string) bool {
	for i := 0; i < len(input); i++ {
		currentSymbol := string(input[i])

		if _, ok := romanSymbolsDecimal[currentSymbol]; !ok {
			return false
		}
	}

	return true
}
