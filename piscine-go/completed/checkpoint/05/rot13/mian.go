package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		input := os.Args[1]

		for _, char := range input {
			rotated := char

			if 'a' <= char && char <= 'z' {
				rotated = 'a' + (char-'a'+13)%26
			} else if 'A' <= char && char <= 'Z' {
				rotated = 'A' + (char-'A'+13)%26
			}

			z01.PrintRune(rotated)
		}

		z01.PrintRune('\n')

	}

}
