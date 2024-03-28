package _1

func ReverseSecondHalf(str string) string {
	result := ""
	if str != "" {
		for i := len(str) - 1; i >= (len(str) / 2); i-- {
			result += string(str[i])
		}
	} else {
		return "Invalid Output\n"
	}
	return result + "\n"

}
