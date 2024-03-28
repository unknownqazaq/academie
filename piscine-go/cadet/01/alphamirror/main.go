package main

import (
	"os"
)

func main() {

	//if len(os.Args) == 2 {
	//	AlphaMirror(os.Args[1])
	//} else {
	//	z01.PrintRune('\n')
	//}

	AlphaMirror("Abc Zyx")

}

func AlphaMirror(str string) {

	word := ""

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			char = 'z' - char + 'a'
			word += string(char)
		} else if char >= 'A' && char <= 'Z' {
			char = 'Z' - char + 'A'
			word += string(char)
		} else {
			word += string(char)
		}
	}
	os.Stdout.WriteString(word + "\n")

}
