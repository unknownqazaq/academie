package main

import (
	student "student/cadet/03"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	f := []func(int, int) int{
		func(accumulator, currentValue int) int {
			return accumulator + currentValue
		},
		func(accumulator, currentValue int) int {
			return accumulator - currentValue
		},
		func(accumulator, currentValue int) int {
			return currentValue * accumulator
		},
	}

	type node struct {
		a         []int
		functions []func(int, int) int
		n         int
	}
	var argInt []int
	table := []node{}

	for i := 0; i < 8; i++ {
		argInt = append(argInt, random.IntSliceBetween(0, 50)...)
		table = append(table, node{
			a:         argInt,
			functions: f,
			n:         random.IntBetween(0, 60),
		})
	}

	table = append(table, node{
		a:         []int{1, 2, 3},
		functions: f,
		n:         93,
	})

	table = append(table, node{
		a:         []int{0},
		functions: f,
		n:         93,
	})

	for _, v := range table {
		for _, f := range v.functions {
			challenge.Function("FoldInt", student.FoldInt, solutions.FoldInt, f, v.a, v.n)
		}
	}
}
