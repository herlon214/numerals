# Numerals
CLI to deal with encoding/decoding of numerals.

## Run
To run you can just type `make run` or run `go run cmd/numerals/main.go`

```
$ make run
 _____ _____ _____ _____ _____ _____ __    _____
|   | |  |  |     |   __| __  |  _  |  |  |   __|
| | | |  |  | | | |   __|    -|     |  |__|__   |
|_|___|_____|_|_|_|_____|__|__|__|__|_____|_____|

Available number systems:
- [0] Decimal
- [1] Binary
- [2] Hexadecimal
- [3] Roman
-------------------------------------------------------------------------
Which one do you want to use? (type the system's number)
```

## Code coverage
```
?   	github.com/herlon214/numerals/cmd/numerals	[no test files]
ok  	github.com/herlon214/numerals/pkg/number	0.247s	coverage: 100.0% of statements in ./pkg/...
github.com/herlon214/numerals/pkg/number/binary.go:11:		NewBinarySystem		100.0%
github.com/herlon214/numerals/pkg/number/binary.go:15:		Name			100.0%
github.com/herlon214/numerals/pkg/number/binary.go:19:		Decode			100.0%
github.com/herlon214/numerals/pkg/number/binary.go:28:		Encode			100.0%
github.com/herlon214/numerals/pkg/number/decimal.go:9:		NewDecimalSystem	100.0%
github.com/herlon214/numerals/pkg/number/decimal.go:13:		Name			100.0%
github.com/herlon214/numerals/pkg/number/decimal.go:17:		Decode			100.0%
github.com/herlon214/numerals/pkg/number/decimal.go:26:		Encode			100.0%
github.com/herlon214/numerals/pkg/number/hexadecimal.go:12:	NewHexadecimalSystem	100.0%
github.com/herlon214/numerals/pkg/number/hexadecimal.go:16:	Name			100.0%
github.com/herlon214/numerals/pkg/number/hexadecimal.go:20:	Decode			100.0%
github.com/herlon214/numerals/pkg/number/hexadecimal.go:29:	Encode			100.0%
github.com/herlon214/numerals/pkg/number/roman.go:36:		NewRomanSystem		100.0%
github.com/herlon214/numerals/pkg/number/roman.go:40:		Name			100.0%
github.com/herlon214/numerals/pkg/number/roman.go:44:		Decode			100.0%
github.com/herlon214/numerals/pkg/number/roman.go:80:		Encode			100.0%
github.com/herlon214/numerals/pkg/number/roman.go:97:		isValid			100.0%
total:								(statements)		100.0%
```

## How the conversions are made
![System's Flow](numerals.png)