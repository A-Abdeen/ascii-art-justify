package asciiart

import "fmt"

func RowPrinterJustify(splitInput []string, alignment string, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	rowlength := FindSize() - 2
	extraspaces := ""

	for _, singleLine := range splitInput { // to print one line at a time
		lengthofwords := 0
		numberofwords := 0
		if singleLine != "" {
			for i := 2; i < 10; i++ { // to print 8 lines of each character
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					if (charRune == 32) && i == 2 {
						numberofwords++
					}
					if i == 2 && charRune != 32 {
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

					if alignment == "--align=justify" && (charRune == ' ') {
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
