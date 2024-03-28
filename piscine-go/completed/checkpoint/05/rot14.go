package _5

func Rot14(s string) string {
	result := ""
	for _, char := range s {

		if 'a' <= char && char <= 'z' {
			result += string((char-'a'+14)%26 + 'a')
		} else if 'A' <= char && char <= 'Z' {
			result += string((char-'A'+14)%26 + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}
