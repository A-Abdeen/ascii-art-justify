package asciiart

import "fmt"

func RowPrinterjustify(splitInput []string, alignment string, sourceFile []byte, f func([]byte, int) []byte) {
	for _, singleLine := range splitInput { // to print one line at a time
		fullRowData := ""
		rowlength := Findsize() - 2
		extraspaces := ""
		lengthofwords := 0
		numberofwords := 0
		if singleLine != "" {
			for i := 2; i < 10; i++ { // to print 8 lines of each character
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					if (charRune == 32) && i == 2 {
						numberofwords++
					}                                          // Calculate number of words and length of words used to find the lengthbetweenwords
					if i == 2 && charRune != 32 {              // to fit the terminal
						lengthofwords += len(charRowData)
					}
				}
			}
			if numberofwords == 0 {
				numberofwords++
			}
			lengthbetweenwords := ((rowlength - lengthofwords) / numberofwords)
			extraspaces = ""
			for i := 0; i < lengthbetweenwords; i++ {
				extraspaces = extraspaces + " "
			}
			for i := 2; i < 10; i++ { // to print 8 lines of each character
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)

					if charRune == ' ' {
						fullRowData = fullRowData + extraspaces
					}
					if charRune != 32 {
						fullRowData = fullRowData + string(charRowData)
					}
				}
				fmt.Println(fullRowData)

				fullRowData = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
}
