package main

import (
	"os"
)

func main() {
	if len(os.Args) == 4 {
		result := searchreplace(os.Args[1], os.Args[2], os.Args[3])
		os.Stdout.WriteString(result)
		os.Stdout.WriteString("\n")
	}
}

func searchreplace(str, a, b string) string {
	var result string

	for _, c := range str {
		if string(c) == a {
			result += b
		} else {
			result += string(c)
		}
	}

	return result
}
