package main

import (
	student "student/completed"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	challenge.Function("PrintComb", student.PrintComb, solutions.PrintComb)
}
