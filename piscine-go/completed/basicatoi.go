package completed

func BasicAtoi(s string) int {
	result := 0
	for _, i := range s {
		if i >= '0' && i <= '9' {
			result = result*10 + int(i-'0')
		} else {
			return 0
		}
	}
	return result
}
