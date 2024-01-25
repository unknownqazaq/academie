package quad

import "github.com/01-edu/z01"

func FirstLineC(x int) {
	z01.PrintRune('A')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('A')
	}
	z01.PrintRune('\n')
}

func LastLineC(x int) {
	z01.PrintRune('C')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
	}
	z01.PrintRune('\n')
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		FirstLineC(x)
		return
	}
	FirstLineC(x)
	if y == 2 {
		LastLineC(x)
		return
	}
	for j := 2; j < y; j++ {
		if j == y {
			LastLineC(x)
			break
		}
		z01.PrintRune('B')
		if x > 1 {
			for i := 2; i < x; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('B')
		}

		z01.PrintRune('\n')
	}
	LastLineC(x)
}

// func main() {
// 	QuadC(5, 5)
// }
