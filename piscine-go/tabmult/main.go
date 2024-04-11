package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		num, err := strconv.Atoi(args)
		if err != nil {
			return
		}
		tabmult(num)

	}

}

func tabmult(num int) {

	for i := 1; i <= 9; i++ {
		sum := i * num

		os.Stdout.WriteString(itoa(i) + " x " + itoa(num) + " = " + itoa(sum) + "\n")
	}

}

func itoa(num int) string {
	result := ""
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num /= 10
	}
	return result

}
