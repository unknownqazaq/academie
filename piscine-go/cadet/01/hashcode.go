package _1

func HashCode(dec string) string {

	str := ""
	for _, char := range dec {
		hash := (int(char) + len(dec)) % 127
		if hash < 32 || hash > 127 {
			hash += 33
		}
		str += string(rune(hash))

	}
	return str

}
