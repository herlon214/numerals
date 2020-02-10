package number

// Encoder contains the function to encode a number input
type Encoder interface {
	Encode(input int64) (string, bool)
}

// Decoder contains the function to decode a number input
type Decoder interface {
	Decode(input string) (int64, bool)
}

// System represents a number system that can encode/decode numbers
type System interface {
	Name() string
	Encoder
	Decoder
}

// Systems returns the list of available number systems
func Systems() []System {
	systems := make([]System, 0)
	systems = append(systems, NewDecimalSystem())
	systems = append(systems, NewBinarySystem())
	systems = append(systems, NewHexadecimalSystem())
	systems = append(systems, NewRomanSystem())

	return systems
}
