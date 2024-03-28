package piscine_go

import "fmt"

func Chunk(slice []int, size int) {

	if size == 0 {
		fmt.Println()
		return

	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return

	}
	arr := [][]int{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		arr = append(arr, slice[i:end])

	}
	fmt.Println(arr)
}
