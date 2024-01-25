package _2

func NRune(s string, n int) rune {
	for index, value := range s { // перевод массив рун в символы (разбиваем на отдельные элементы )
		if index == n-1 {
			return value
		}
	}
	return 0
}
