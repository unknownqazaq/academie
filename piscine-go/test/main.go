package main

import (
	"fmt"
	piscine_go "student"
)

func main() {
	fmt.Println(piscine_go.SwapLast([]int{1, 2, 3, 4}))
	fmt.Println(piscine_go.SwapLast([]int{3, 4, 5}))
	fmt.Println(piscine_go.SwapLast([]int{1}))
}
