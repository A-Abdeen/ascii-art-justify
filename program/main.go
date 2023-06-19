package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	// Check if input is correct
	if len(os.Args) != 4 {
		fmt.Print("\nCorrect format: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard\n\n")
		return
	}
	alignment := os.Args[1]
	rawInput := os.Args[2]
	file := os.Args[3]
	switch {
	case file == "standard":
		file = "standard.txt"
	case file == "shadow":
		file = "shadow.txt"
	case file == "thinkertoy":
		file = "thinkertoy.txt"
	default:
		fmt.Println("\nAvailable banner formats are: standard, shadow or thinkertoy.")
		return
	}
	sourceFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	asciiart.RowPrinter(splitInput, alignment, sourceFile, asciiart.RowParser)
}
