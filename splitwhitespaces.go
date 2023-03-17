package piscine

func SplitWhiteSpaces(s string) []string {
	var substrings []string
	word := ""
	for i, c := range s {
		if c != ' ' && c != '\n' && c != '\t' {
			word += string(c)
		}
		if ((c == ' ' || c == '\t' || c == '\n') && word != "") || i == len(s)-1 {
			substrings = append(substrings, word)
			word = ""
		}
	}
	return substrings
}

// func main() {
// 	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
// }
