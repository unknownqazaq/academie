package main

import "os"

/*
expandstr
Instructions
Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

The string will be followed by a newline ('\n').

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.

Usage
$ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
$ go run . "   only  it's harder   " | cat -e
only   it's   harder$
$ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
$ go run .
$
*/
func main() {
	//if len(os.Args) == 2 {
	//	expandstr(os.Args[1])
	//
	//}
	expandstr("   only  it's harder   ")

}

func expandstr(str string) {
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
	for i, word := range words {
		if i != 0 {
			os.Stdout.WriteString("   ")
		}
		os.Stdout.WriteString(word)
	}
	os.Stdout.WriteString("\n")
}
