package main

import (
	"fmt"
	piscine_go "student"
)

func main() {
	fmt.Println(piscine_go.AtoiBase("125", "0123456789"))
	fmt.Println(piscine_go.AtoiBase("1111101", "01"))
	fmt.Println(piscine_go.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine_go.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine_go.AtoiBase("bbbbbab", "-ab"))
}
