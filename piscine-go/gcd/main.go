package main

import (
	"fmt"
	"os"
)

/*
allowedFunctions
:
["fmt.*", "os.Args", "strconv.Atoi", "--allow-builtin"]
gcd
Instructions
Write a program that takes two string representing two strictly positive integers that fit in an int.

The program displays their greatest common divisor followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing.

All arguments tested will be positive int values.

Usage
$ go run . 42 10 | cat -e
2$
$ go run . 42 12
6
$ go run . 14 77
7
$ go run . 17 3
1
$ go run .
$ go run . 50 12 4
$
*/
func main() {
	if len(os.Args) == 3 {
		fmt.Println(gcd(os.Args[1], os.Args[2]))

	}

}

func gcd(str1, str2 string) int {
	num1 := atoi(str1)
	num2 := atoi(str2)
	max := num1

	if num1 < num2 {
		max = num2

	}
	for i := max; i > 0; i-- {
		if num1%i == 0 && num2%i == 0 {
			return i
		}
	}

	return 0

}

func atoi(str string) int {
	result := 0
	sign := 1

	for i, char := range str {
		if i == 0 && (char == '+' || char == '-') {
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
