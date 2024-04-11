package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println()
	} else {
		for _, v := range os.Args[1:] {
			if brackets(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	}
}

func brackets(s string) bool {
	var stack []rune

	for _, char := range s {
		if isOpeningBracket(char) {
			stack = append(stack, char)
		} else if isClosingBracket(char) {
			if len(stack) == 0 || !areMatchingBrackets(stack[len(stack)-1], char) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isOpeningBracket(char rune) bool {
	return char == '(' || char == '[' || char == '{'
}

func isClosingBracket(char rune) bool {
	return char == ')' || char == ']' || char == '}'
}

func areMatchingBrackets(opening, closing rune) bool {
	switch opening {
	case '(':
		return closing == ')'
	case '[':
		return closing == ']'
	case '{':
		return closing == '}'
	default:
		return false
	}
}
