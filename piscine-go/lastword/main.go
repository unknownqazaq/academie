package main

import (
	"os"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		if args == " " {
			return

		}
		result := lastword(args)
		os.Stdout.WriteString(result + "\n")
	}

}

func lastword(str string) string {
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
	return words[len(words)-1]

}
