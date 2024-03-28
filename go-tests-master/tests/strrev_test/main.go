package main

import (
	student "student/completed/checkpoint/02"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		random.StrSlice(chars.ASCII),
		"Hello!",
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		challenge.Function("StrRev", student.StrRev, solutions.StrRev, arg)
	}
}
