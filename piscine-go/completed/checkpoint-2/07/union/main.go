package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n') // Print a newline if the number of arguments is not 2
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	seen := make(map[rune]bool)

	for _, char := range str1 {
		if !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}

	for _, char := range str2 {
		if !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}

	z01.PrintRune('\n') // Print a newline after processing the strings
}
