package quad

import "github.com/01-edu/z01"

func FirstLineE(x int) {
	z01.PrintRune('A')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
	}
	z01.PrintRune('\n')
}

func LastLineE(x int) {
	z01.PrintRune('C')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('A')
	}
	z01.PrintRune('\n')
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		FirstLineE(x)
		return
	}
	FirstLineE(x)
	if y == 2 {
		LastLineE(x)
		return
	}
	for j := 2; j < y; j++ {
		if j == y {
			LastLineE(x)
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
	LastLineE(x)
}

// func main() {
// 	QuadE(1, 5)
// }
