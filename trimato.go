package asciiart

func TrimAto(s string) int {
	runes := []rune(s)
	var n int = 0
	length := len(runes)
	for i := 0; i < length; i++ {
		if s[i] <= '9' && s[i] >= '0' {
			n = (n*10 + int(s[i]) - '0')
		}
	}
	for i := 0; i < length; i++ {
		if s[i] <= '9' && s[i] >= '0' {
			break
		}
		if s[i] == 45 {
			n = n * -1
			return n
		}
	}
	return n
}
