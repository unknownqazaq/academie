package main

import (
	"github.com/01-edu/z01"
	"os"
)

func searchReplace(str, search, replace string) {

	for _, char := range str {
		if string(char) == search {
			z01.PrintRune(rune(replace[0]))
		} else {
			z01.PrintRune(char)
		}
	}
}

func main() {
	if len(os.Args) == 4 {
		searchReplace(os.Args[1], os.Args[2], os.Args[3])
		z01.PrintRune('\n')
	}
}
