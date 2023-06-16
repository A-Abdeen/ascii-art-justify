package asciiart

import "fmt"

func RowPrinter(splitInput []string, alignment string, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	rowlength := Findsize() - 2
	rowlength2 := 1
	
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
			for i := 2; i < 10; i++ { // to print 8 lines of each character
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					fullRowData = fullRowData + string(charRowData)
				}
				if alignment == "--align=left" {
					rowlength2 = rowlength - len(fullRowData)
				fmt.Print("|")
					fmt.Print(fullRowData)

					fullRowData = ""
					for j := 0; j < rowlength2; j++ {
						fmt.Print(" ")
					}
					fmt.Println("|")
				}
				if alignment == "--align=right" {
					fmt.Print("|")
					rowlength2 = rowlength - len(fullRowData)
				for j := 0; j < rowlength2; j++ {
					fmt.Print(" ")
				}
				fmt.Print(fullRowData)
				fullRowData = ""
				fmt.Println("|")
			}
			 if alignment == "--align=justify" {
				finalinput := ""
				extraspaces := ""
		lengthofwords := 0
		count := 0
		for _, word := range splitInput {
			lengthofwords = len(word) + lengthofwords
			count++
		}
		lengthbetweenwords := (rowlength - lengthofwords) / count
		for _, word := range splitInput {
			for i := 0; i < lengthbetweenwords; i++ {
				extraspaces = extraspaces + " "
			}
			finalinput = finalinput + extraspaces + word
			extraspaces = ""
		}
			 }  	
		if alignment == "--align=center" {
			fmt.Print("|")
			rowlength2 = (rowlength - len(fullRowData)) / 2
		for j := 0; j < rowlength2; j++ {
			fmt.Print(" ")
		}
		fmt.Print(fullRowData)
		fullRowData = ""
		for j := 0; j < rowlength2; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	} 
	}
}else {
	fmt.Print("\n")}
}
}
