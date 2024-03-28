package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		result := union(os.Args[1], os.Args[2])
		os.Stdout.WriteString(result + "\n")

	} else {
		os.Stdout.WriteString("\n")
	}

}
func union(str1, str2 string) string {
	word := str1 + str2
	arrbool := make(map[rune]bool)
	result := ""
	for _, char := range word {
		if !arrbool[char] {
			result += string(char)
			arrbool[char] = true

		}

	}
	return result

}
