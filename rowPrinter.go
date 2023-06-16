package asciiart

import "fmt"

func RowPrinter(splitInput []string, alignment string, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	rowlength := Findsize() - 2
	rowlength2 := 1
	extraspaces := ""
	lengthofwords := 0
	numberofwords := 0
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
			for i := 2; i < 10; i++ { // to print 8 lines of each character
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
				if alignment == "--align=justify" && (charRune == 32 || charRune == 10 ) && i == 2 {
						numberofwords++
					}
					if alignment == "--align=justify" && i == 2 && charRune != 32 {
						lengthofwords += len(charRowData)
					}
					fullRowData = fullRowData + string(charRowData)
				}

				if alignment == "--align=left" { // to print if alignment =left
					rowlength2 = rowlength - len(fullRowData)
					fmt.Println(fullRowData)
					fullRowData = ""
				}
				if alignment == "--align=right" { // to print if alignment =right
					rowlength2 = rowlength - len(fullRowData)
					for j := 0; j < rowlength2; j++ {
						fmt.Print(" ")
					}
					fmt.Println(fullRowData)
					fullRowData = ""
				}
				if alignment == "--align=center" { // to print if alignment = center
					rowlength2 = (rowlength - len(fullRowData)) / 2
					for j := 0; j < rowlength2; j++ {
						fmt.Print(" ")
					}
					fmt.Print(fullRowData)
					fullRowData = ""
					for j := 0; j < rowlength2; j++ {
						fmt.Print(" ")
					}
					fmt.Print("\n")
				}
			}
		} else {
			fmt.Print("\n")
		}
	}
	if alignment == "--align=justify" { // to print if alignment = justify
		fullRowData = ""
		if numberofwords == 0 {
			numberofwords++
		}
		lengthbetweenwords := ((rowlength - lengthofwords) / numberofwords)
		for i := 0; i < lengthbetweenwords; i++ {
			extraspaces = extraspaces + " "
		}
		for _, singleLine := range splitInput { // to print one line at a time
			if singleLine != "" {
				for i := 2; i < 10; i++ { // to print 8 lines of each character
					for _, charRune := range singleLine { // to combine the prespective line of each char to the next
						rowLocation := (((int(charRune) - 32) * 9) + i)
						charRowData := f(sourceFile, rowLocation)

						if alignment == "--align=justify" && (charRune == ' ' || charRune == 10) {
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
}
