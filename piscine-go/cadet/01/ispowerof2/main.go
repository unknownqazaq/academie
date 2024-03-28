package main

import "os"

func main() {
	if len(os.Args) == 2 {
		if IsPowerOf2(os.Args[1]) {
			os.Stdout.WriteString("true" + "\n")
		} else {
			os.Stdout.WriteString("false" + "\n")
		}
	}
}

func IsPowerOf2(args string) bool {
	a := atoi(args)
	return a > 0 && (a&(a-1)) == 0
}

func atoi(str string) int {
	result := 0
	for _, char := range str {
		if '0' <= char || char >= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return result
}
