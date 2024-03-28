package _1

func DigitLen(n, base int) int {
	count := 0

	if base >= 2 && base <= 36 {

		for n != 0 {
			n = n / base
			count++
		}
	} else {
		return -1
	}

	return count

}
