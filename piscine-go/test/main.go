package main

import (
	"github.com/01-edu/z01"
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
	searchReplace("hella there", "a", "o")

}
