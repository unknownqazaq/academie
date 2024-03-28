package _1

import "github.com/01-edu/z01"

func ReduceInt(a []int, f func(int, int) int) {

	result := 0
	for i := 1; i < len(a); i++ {
		result = f(a[i-1], a[i])
	}
	str := Itoa(result)
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}

func Itoa(num int) string {
	result := ""
	sign := ""
	if num < 0 {
		num *= -1
		sign = "-"
	}
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num /= 10

	}
	return sign + result
}
