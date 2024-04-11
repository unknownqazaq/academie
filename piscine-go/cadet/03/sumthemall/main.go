package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if there are enough arguments provided
	if len(os.Args) < 2 {
		fmt.Println("0")
		return
	}

	sum := 0
	// Iterate over the command-line arguments starting from index 1
	for _, arg := range os.Args[1:] {
		// Convert the argument to an integer
		num, err := strconv.Atoi(arg)
		if err != nil {
			// If conversion fails (e.g., argument is not a number), print 0 and exit
			fmt.Println("0")
			return
		}
		// Check for integer overflow
		if (sum > 0 && num > (1<<31-1)-sum) || (sum < 0 && num < (-1<<31)-sum) {
			fmt.Println("0")
			return
		}
		// Update the sum
		sum += num
	}

	// Print the sum
	fmt.Println(sum)
}
