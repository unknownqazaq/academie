package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	// Check if there is at least one command-line argument
	if len(os.Args) > 1 {
		// Display the first command-line argument
		for _, i2 := range os.Args[1] {
			z01.PrintRune(i2)
		}
		z01.PrintRune('\n')
	}
}
