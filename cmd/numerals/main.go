package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/herlon214/numerals/pkg/number"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	printHeader()

	// Load available systems
	numberSystems := NumberSystems()

	// List available systems
	fmt.Println("Available number systems:")
	for i, system := range numberSystems {
		fmt.Printf("- [%d] %s\n", i, system.Name())
	}

	printSeparator()

	// Ask user to specify one system
	fmt.Println("Which one do you want to use? (type the system's number)")
	input := readLine(reader, "\n")

	// Convert the input to number and check if it's valid
	systemIndex, err := strconv.ParseInt(input, 10, 64)
	if err != nil || int(systemIndex) > len(numberSystems) || int(systemIndex) < 0 {
		fmt.Println("Sorry, I didn't find the specified system.")
		fmt.Println(err.Error())
		return
	}

	printSeparator()

	// Set the input system and ask for the input number
	inputSystem := numberSystems[systemIndex]
	fmt.Printf("Okay, let's use the %s then...\n", inputSystem.Name())
	fmt.Printf("Type your number in the %s system:\n", strings.ToLower(inputSystem.Name()))
	specifiedNumber := readLine(reader, "\n")

	printSeparator()

	// Ask the user to specify the output system
	fmt.Println("To which system do you want to convert? (type the system's number)")
	input = readLine(reader, "\n")

	// Convert the input to number and check if it's valid
	systemIndex, err = strconv.ParseInt(input, 10, 64)
	if err != nil || int(systemIndex) > len(numberSystems) || int(systemIndex) < 0 {
		fmt.Println("Sorry, I didn't find the specified system.")
		fmt.Println(err.Error())
		return
	}

	outputSystem := numberSystems[systemIndex]
	decimalInput, success := inputSystem.Decode(specifiedNumber)
	if !success {
		fmt.Println("Sorry, something went wrong during the number decode process. Check your number and try again...")
		return
	}

	result, success := outputSystem.Encode(decimalInput)
	if !success {
		fmt.Println("Sorry, something went wrong during the number encode process. Check your number and try again...")
		return
	}
	fmt.Printf("The result is: %s\n", result)

	printSeparator()
	fmt.Println("See you!")

	return
}

// NumberSystems returns the list of available number systems
func NumberSystems() []number.System {
	systems := make([]number.System, 0)
	systems = append(systems, number.NewDecimalSystem())
	systems = append(systems, number.NewBinarySystem())
	systems = append(systems, number.NewHexadecimalSystem())
	systems = append(systems, number.NewRomanSystem())

	return systems
}

func printHeader() {
	fmt.Println(` _____ _____ _____ _____ _____ _____ __    _____ 
|   | |  |  |     |   __| __  |  _  |  |  |   __|
| | | |  |  | | | |   __|    -|     |  |__|__   |
|_|___|_____|_|_|_|_____|__|__|__|__|_____|_____|
																									`)
}

func printSeparator() {
	fmt.Println("-------------------------------------------------------------------------")
}

func readLine(reader *bufio.Reader, delim string) string {
	output, err := reader.ReadString(delim[0])
	if err != nil {
		fmt.Printf("Sorry, something went wrong: %s", err.Error())
		os.Exit(1)
	}

	return strings.TrimSuffix(output, delim)
}
