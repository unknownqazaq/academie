package _1

func IsSorted(f func(a, b int) int, a []int) bool {
	firstbool := true
	secandbool := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			firstbool = false

		}
		if f(a[i-1], a[i]) < 0 {
			secandbool = false

		}

	}
	return secandbool || firstbool

}
