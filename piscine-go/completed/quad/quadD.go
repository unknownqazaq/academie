package quad

import "github.com/01-edu/z01"

func FirstLineD(x int) {
	z01.PrintRune('A')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
	}
	z01.PrintRune('\n')
}

func LastLineD(x int) {
	z01.PrintRune('A')
	if x > 1 {
		for i := 2; i < x; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
	}
	z01.PrintRune('\n')
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		FirstLineD(x)
		return
	}
	FirstLineD(x)
	if y == 2 {
		LastLineD(x)
		return
	}
	for j := 2; j < y; j++ {
		if j == y {
			LastLineD(x)
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
	LastLineD(x)
}

// func main() {
// 	QuadD(1, 5)
// }
