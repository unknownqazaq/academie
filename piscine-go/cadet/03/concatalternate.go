package _3

/*
concatalternate
Instructions
Write a function ConcatAlternate() that receives two slices of an int as arguments and returns a new slice with the result of the alternated values of each slice.

The input slices can be of different lengths.
The new slice should start with an element of the largest slice.
If the slices are of equal length, the new slice should return the elements of the first slice first and then the elements of the second slice.
Expected function
func ConcatAlternate(slice1,slice2 []int) []int {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
And its output:

$ go run .
[1 4 2 5 3 6]
[1 2 3 4 5 6 7 8 9 10 11]
[4 1 5 2 6 3 7 8 9]
[1 2 3]
*/

func ConcatAlternate(slice1, slice2 []int) []int {
	var arr []int
	lenSlice1 := len(slice1)
	lenSlice2 := len(slice2)
	max := lenSlice1
	if max < lenSlice2 {
		max = lenSlice2

		for i := 0; i < max; i++ {
			if i%2 == 0 {
				arr = append(arr, slice2[i])

			} else {
				arr = append(arr, slice1[i])
			}

		}

	} else {
		for i := 0; i < max; i++ {
			if i%2 == 0 {
				arr = append(arr, slice1[i])

			} else {
				arr = append(arr, slice2[i])
			}
		}

	}
	return arr

}
