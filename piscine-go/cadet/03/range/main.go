package main

import (
	"fmt"
)

func main() {
	//if len(os.Args) == 3 {
	//	args := os.Args[1]
	//	args2 := os.Args[2]
	//	num, err := strconv.Atoi(args)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	num2, err := strconv.Atoi(args2)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	Range(num, num2)
	//
	//}
	Range(7, -3)

}

func Range(n, n2 int) {
	if n2 < n {
		for i := n; i >= n2; i-- {
			fmt.Print(i)
			if i > n2 {
				fmt.Print(" ")
			}
		}
	} else {
		for i := n; i <= n2; i++ {
			fmt.Print(i)
			if i < n2 {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}
