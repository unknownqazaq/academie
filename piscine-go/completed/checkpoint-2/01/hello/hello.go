package main

import "github.com/01-edu/z01"

func main() {
	hello := "Hello World!"
	for _, char := range hello {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
