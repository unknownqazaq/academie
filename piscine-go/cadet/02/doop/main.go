package main

import (
	"os"
)

func main() {
	if len(os.Args) == 4 {
		doop(os.Args[1], os.Args[2], os.Args[3])
	}

}

func doop(a, operator, b string) {
	num1, boola := atoi(a)
	num2, boolb := atoi(b)
	if !boola || !boolb {
		return
	}
	if int64(num1)+int64(num2) > 1<<31-1 || int64(num1)+int64(num2) < -1<<31 {
		return
	}
	switch operator {
	case "+":
		os.Stdout.WriteString(itoa(num1+num2) + "\n")
	case "-":
		os.Stdout.WriteString(itoa(num1-num2) + "\n")
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		os.Stdout.WriteString(itoa(num1/num2) + "\n")
	case "*":
		os.Stdout.WriteString(itoa(num1*num2) + "\n")
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		os.Stdout.WriteString(itoa(num1%num2) + "\n")
	}
}

func atoi(str string) (int, bool) {
	result := 0
	sign := 1

	for i, char := range str {
		if i == 0 && (char == '+' || char == '-') {
			if char == '-' {
				sign = -1
			}
		} else if char < '0' || char > '9' {
			return 0, false
		} else {
			result = result*10 + int(char-'0')
		}
	}
	return result * sign, true

}

func itoa(num int) string {
	result := ""
	sign := ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = num * -1
		sign = "-"
	}
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num /= 10
	}
	return sign + result
}
