package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		for _, char := range arg {
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
				repeats := 0
				if char >= 'a' && char <= 'z' {
					repeats = int(char - 'a' + 1)
				} else {
					repeats = int(char - 'A' + 1)
				}
				for i := 0; i < repeats; i++ {
					z01.PrintRune(char)
				}
			} else {
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')

	}
}
