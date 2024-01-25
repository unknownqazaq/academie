package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		s := os.Args[1]
		for _, char := range s {
			if 'a' <= char && char <= 'z' {
				z01.PrintRune(char - 32)
			} else if 'A' <= char && char <= 'Z' {
				z01.PrintRune(char + 32)
			} else {
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')

	}
}
