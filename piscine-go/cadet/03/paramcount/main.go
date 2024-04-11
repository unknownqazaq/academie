package main

import "os"

func main() {
	args := os.Args[1:]
	res := itoa(paramcount(args))
	os.Stdout.WriteString(res + "\n")

}

func paramcount(str []string) int {
	return len(str)
}

func itoa(num int) string {
	if num == 0 {
		return "0"

	}
	result := ""

	if num > 0 {
		digin := num % 10
		result = string(rune(digin+'0')) + result
		num /= 10
	}
	return result

}
