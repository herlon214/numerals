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
