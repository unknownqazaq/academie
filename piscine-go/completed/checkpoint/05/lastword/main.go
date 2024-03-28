package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		lastword := ""
		for i := len(args) - 1; i >= 0; i-- {
			if args[i] != ' ' {
				lastword = string(args[i]) + lastword
			} else if lastword != "" {
				break
			}
		}
		if lastword != "" {
			for _, char := range lastword {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')

		}
	}
}
