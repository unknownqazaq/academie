package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		cleanstr(os.Args[1])
	} else {
		fmt.Println()
		return
	}
}

func cleanstr(str string) {
	word := ""
	words := []string{}

	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	fmt.Println(join(words))

}

func join(words []string) string {
	str := ""
	for i, word := range words {
		str += word
		if i != len(words)-1 {
			str += " "
		}
	}
	return str
}
