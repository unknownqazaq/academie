package _3

func Itoa(n int) string {
	sign := ""
	result := ""
	if n == 0 {
		return "0"

	}
	if n < 0 {
		n *= -1
		sign = "-"
	}

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10

	}
	return sign + result

}
