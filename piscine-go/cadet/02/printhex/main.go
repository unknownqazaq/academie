package main

import (
	"os"
)

/*
printhex
Instructions
Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline ('\n').

If the number of arguments is different from 1, the program displays nothing.
Error cases have to be handled as shown in the example below.
Usage
$ go run . 10
a
$ go run . 255
ff
$ go run . 5156454
4eae66
$ go run .
$ go run . "123 132 1" | cat -e
ERROR$
$
*/

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		num := atoi(args)
		printhex(num)

	}

}

func printhex(n int) {
	hexChars := "0123456789abcdef"
	result := ""
	for n > 0 {
		index := n % 16
		result = string(hexChars[index]) + result
		n /= 16
	}
	if result == "" {
		result = "0"
	}
	os.Stdout.WriteString(result + "\n")
}

func atoi(str string) int {
	result := 0
	sign := 1

	for i, char := range str {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				sign = -1

			}
		} else if char < '0' || char > '9' {
			return 0
		} else {
			result = result*10 + int(char-'0')
		}
	}
	return result * sign

}
