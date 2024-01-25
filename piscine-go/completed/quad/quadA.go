package quad

import "github.com/01-edu/z01"

func FirstLineA(x int) {
	if x == 1 {
		z01.PrintRune('o')
	} else {
		z01.PrintRune('o')
		for i := 2; i < x; i++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
	}
	z01.PrintRune('\n')
}

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		FirstLineA(x)
		return
	}
	FirstLineA(x)
	for j := 2; j < y; j++ {
		if x == 1 {
			z01.PrintRune('|')
		} else {
			z01.PrintRune('|')
			for i := 2; i < x; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('|')
		}
		z01.PrintRune('\n')
	}
	FirstLineA(x)
}

// func main() {
// 	QuadA(1, 5)
// }
