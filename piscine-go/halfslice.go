package piscine_go

/*


halfslice
Instructions
Write a function HalfSlice() that receives a slice of int and returns another slice of int with the first half of the values. If the length is odd, round it up.

Expected function
func HalfSlice(slice []int) []int {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
}
And its output:

$ go run . | cat -e
[1 2 3 4 5]
[1 2]
*/

func HalfSlice(slice []int) []int {
	var arr []int
	lenSlice := len(slice)
	if lenSlice == 0 {
		return arr
	}
	if lenSlice%2 != 0 {
		lenSlice += 1
	}

	for i := 0; i < lenSlice/2; i++ {
		arr = append(arr, slice[i])
	}
	return arr
}
