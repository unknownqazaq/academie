package _2

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)

	}
}
