package main

import (
	"sort"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func isSortedBy1(a, b int) int {
	if a-b < 0 {
		return -1
	}
	if a-b > 0 {
		return 1
	}
	return 0
}

func isSortedBy10(a, b int) int {
	if a-b < 0 {
		return -10
	}
	if a-b > 0 {
		return 10
	}
	return 0
}

func isSortedByDiff(a, b int) int {
	return a - b
}

func isSortedWrong(a, b int) int {
	return -1
}

func main() {
	functions := []func(int, int) int{isSortedByDiff, isSortedBy1, isSortedBy10, isSortedWrong}

	type node struct {
		f func(int, int) int
		a []int
	}

	table := []node{}

	// 5 unordered slices
	for i := 0; i < 5; i++ {
		function := functions[random.IntBetween(0, len(functions)-1)]
		val := node{
			f: function,
			a: random.IntSliceBetween(-1000000, 1000000),
		}
		table = append(table, val)
	}

	// 5 slices ordered in ascending order
	for i := 0; i < 5; i++ {
		function := functions[random.IntBetween(0, len(functions)-1)]
		ordered := random.IntSliceBetween(-1000000, 1000000)
		sort.Ints(ordered)

		val := node{
			f: function,
			a: ordered,
		}
		table = append(table, val)
	}

	// 5 slices ordered in descending order
	for i := 0; i < 5; i++ {
		function := functions[random.IntBetween(0, len(functions)-1)]
		reversed := random.IntSliceBetween(-1000000, 1000000)
		sort.Sort(sort.Reverse(sort.IntSlice(reversed)))
		val := node{
			f: function,
			a: reversed,
		}
		table = append(table, val)
	}

	table = append(table, node{
		f: isSortedByDiff,
		a: []int{1, 2, 3, 4, 5, 6},
	})

	table = append(table, node{
		f: isSortedByDiff,
		a: []int{6, 5, 4, 3, 2, 1},
	})

	table = append(table, node{
		f: isSortedByDiff,
		a: []int{0, 0, 0, 0, 0, 0, 0},
	})

	table = append(table, node{
		f: isSortedByDiff,
		a: []int{0},
	})

	for _, arg := range table {
		challenge.Function("IsSorted", student.IsSorted, solutions.IsSorted, arg.f, arg.a)
	}
}
