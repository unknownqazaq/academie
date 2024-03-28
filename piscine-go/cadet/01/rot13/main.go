package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		rot13(os.Args[1])

	}

}

func rot13(str string) {

	for _, char := range str {
		result := char
		if char >= 'a' && char <= 'z' {
			result = 'a' + (char-'a'+13)%26
		} else if char >= 'A' && char <= 'Z' {
			result = 'A' + (char-'A'+13)%26
		}

		z01.PrintRune(result)

	}
	z01.PrintRune('\n')
}
