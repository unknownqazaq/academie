package _1

func SetSpace(s string) string {
	if s == "" {
		return ""
	}
	if s == " " {
		return "Error"

	}
	if s[0] < 'A' || s[0] > 'Z' {
		return "Error"
	}

	word := ""
	for i, char := range s {
		if i == 0 {
			word += string(char)
			continue
		} else if char >= 'a' && char <= 'z' {
			word += string(char)
		} else if char >= 'A' && char <= 'Z' {
			word += " " + string(char)
		} else {
			return "Error"
		}

	}
	return word
}
