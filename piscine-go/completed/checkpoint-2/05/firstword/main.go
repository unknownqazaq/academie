package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		if args != "" {
			flag := true
			for _, char := range args {
				if char != ' ' {
					z01.PrintRune(char)
					flag = false
				} else if !flag {
					break
				}
			}
			z01.PrintRune('\n')

		}

	}

}
