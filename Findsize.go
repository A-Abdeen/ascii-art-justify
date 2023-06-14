package asciiart


import (
	"os"
	"log"
	"os/exec"
)

func Findsize() int{
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	rowlength := 1
	out, err := cmd.Output()
	for i := 0; i < len(out); i++ {
		if string(out[i]) == " " {
			rowlength = TrimAto(string(out[i:]))
			break
		} 
	}
	if err != nil {
	  log.Fatal(err)
	}
	return rowlength
}