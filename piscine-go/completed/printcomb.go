package completed

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				z01.PrintRune(rune(i) + '0')
				z01.PrintRune(rune(j) + '0')
				z01.PrintRune(rune(k) + '0')
				if i+j+k == 24 {
					break
				}
				z01.PrintRune(',')
				z01.PrintRune(' ')

			}

		}

	}
	z01.PrintRune('\n')

}
