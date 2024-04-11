package piscine_go

// Функция AtoiBase принимает строку числа и строку базы и возвращает целочисленное значение числа в заданной базе.
func AtoiBase(s string, base string) int {
	if len(s) < 2 {
		return 0

	}
	for _, char := range base {
		if char == '-' || char == '+' {
			return 0
		}
	}
	baseMap := make(map[rune]int)

	for i, char := range base {
		if _, exists := baseMap[char]; exists {
			return 0

		}
		baseMap[char] = i
	}
	result := 0
	baseLen := len(base)
	for _, char := range s {
		if _, exists := baseMap[char]; !exists {
			return 0
		}
		result = result*baseLen + baseMap[char]

	}
	return result
}
