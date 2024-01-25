package quad

import "github.com/01-edu/z01"

func FirstLineB(x int) {
	z01.PrintRune('/')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('\\')
	}
	z01.PrintRune('\n')
}

func LastLineB(x int) {
	z01.PrintRune('\\')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('/')
	}
	z01.PrintRune('\n')
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		FirstLineB(x)
		return
	}
	FirstLineB(x)
	if y == 2 {
		LastLineB(x)
		return
	}
	for j := 2; j < y; j++ {
		if j == y {
			LastLineB(x)
		}
		z01.PrintRune('*')
		if x > 1 {
			for i := 2; i < x; i++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('*')
		}

		z01.PrintRune('\n')
	}
	LastLineB(x)
}

// func main() {
// 	QuadB(5, 5)
// }
