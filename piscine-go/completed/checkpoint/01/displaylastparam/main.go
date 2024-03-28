package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	lens := len(os.Args)
	// Check if there is at least one command-line argument
	if lens > 1 {
		// Display the first command-line argument
		for _, i2 := range os.Args[lens-1] {
			z01.PrintRune(i2)
		}
		z01.PrintRune('\n')
	}
}
