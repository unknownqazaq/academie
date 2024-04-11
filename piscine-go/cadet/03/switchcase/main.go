package main

import "os"

func main() {
	if len(os.Args) == 2 {
		result := switchcase(os.Args[1])
		os.Stdout.WriteString(result + "\n")
	}
}

func switchcase(str string) string {
	word := ""
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			word += string(char + 32)
		} else if char >= 'a' && char <= 'z' {
			word += string(char - 32)
		} else {
			word += string(char)
		}
	}
	return word
}
