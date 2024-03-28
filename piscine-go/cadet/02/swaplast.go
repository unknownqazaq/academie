package piscine_go

func SwapLast(slice []int) []int {
	if len(slice) > 1 {
		slice[len(slice)-1], slice[len(slice)-2] = slice[len(slice)-2], slice[len(slice)-1]
		return slice
	}
	return slice
}
