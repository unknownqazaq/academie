package main

import "os"

func main() {
	if len(os.Args) == 2 {
		num := atoi(os.Args[1])
		result := addprimesum(num)
		str := itoa(result)
		os.Stdout.WriteString(str + "\n")

	} else {
		os.Stdout.WriteString("0\n")
	}

}
func addprimesum(num int) int {
	sum := 0
	if num <= 0 {
		return 0

	}
	for i := num; i > 1; i-- {
		if isprime(i) {
			sum += i
		}

	}
	return sum

}

func isprime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

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

func itoa(num int) string {
	result := ""
	sign := ""
	if num == 0 {
		return "0"

	} else if num < 0 {
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
