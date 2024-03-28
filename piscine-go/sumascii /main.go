package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		sunascii(os.Args[1])

	} else {
		z01.PrintRune('0')
	}

}

func sunascii(str string) int {
	arrByte := []byte(str)
	var result byte
	for _, char := range arrByte {
		result += char
	}
	return result

}
