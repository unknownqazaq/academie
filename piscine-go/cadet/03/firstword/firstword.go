package main

import "os"

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		if args == "" {
			return
		}
		os.Stdout.WriteString(firstword(args))
	}
	os.Stdout.WriteString("\n")

}
func firstword(str string) string {
	word := ""
	words := []string{}
	for _, char := range str {
		if char == ' ' && word != "" {
			words = append(words, word)
			word = ""
		} else if char != ' ' {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words[0]

}
