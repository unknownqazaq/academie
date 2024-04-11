package _3

// Split - функция принимает строку 's' и разделитель 'sep' и разбивает строку
// на части там, где встречается разделитель. Она возвращает срез, содержащий
// отдельные части строки.
func Split(s, sep string) []string {
	// Создаем пустой срез для хранения частей строки после разделения.
	var result []string

	// 'start' будет отслеживать начало следующей части строки.
	start := 0

	// Проходим по каждому символу в строке 's'.
	for i := 0; i < len(s); i++ {
		// Проверяем, соответствует ли текущая подстрока, начиная с 'i', разделителю 'sep'.
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			// Если соответствует, значит, мы нашли разделитель.
			// Извлекаем часть строки от 'start' до 'i' (исключая разделитель).
			index := s[start:i]

			// Добавляем эту часть в срез результата.
			result = append(result, index)

			// Обновляем 'start' до индекса сразу после разделителя.
			start = i + len(sep)

			// Перемещаем 'i' на 'start - 1', чтобы на следующей итерации
			// 'i' увеличилось до 'start', что позволяет пропустить символы,
			// которые были частью разделителя.
			i = start - 1
		}
	}

	// После завершения цикла добавляем оставшуюся часть строки (после последнего разделителя)
	// в срез результата.
	result = append(result, s[start:])

	// Наконец, возвращаем срез, содержащий отдельные части строки.
	return result
}
