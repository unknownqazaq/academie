package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
nstructions
Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

Factors must be displayed in ascending order and separated by *.

If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.


*/

func main() {
	if len(os.Args) == 2 {
		args, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		fprime(args)
	}

}

func fprime(num int) {
	if num <= 1 {
		return
	}
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(i)
			num /= i
			i--
			if i < num {
				fmt.Print("*")
			}
		}
	}
	fmt.Println()
}
