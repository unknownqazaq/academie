package _1

func NotDecimal(dec string) string {

	for _, char := range dec {
		if char >= 'A' && char <= 'Z' {
			return dec + "\n"
		} else if char >= 'a' && char <= 'z' {
			return dec + "\n"
		} else if char == '+' {
			return dec + "\n"
		}
	}
	not := atoi(dec)
	return itoa(not) + "\n"
}

func atoi(str string) int {
	result := 0
	sign := 1

	for i, char := range str {
		if i == 0 && char == '-' {
			sign *= -1
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else if char == '.' {
			continue
		}
	}
	return result * sign
}

func itoa(num int) string {
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
