package main

import (
	"fmt"
	piscine_go "student"
)

func main() {
	fmt.Println(piscine_go.IsCapitalized("Hello! How are you?"))
	fmt.Println(piscine_go.IsCapitalized("Hello How Are You"))
	fmt.Println(piscine_go.IsCapitalized("Whats 4this 100K?"))
	fmt.Println(piscine_go.IsCapitalized("Whatsthis4"))
	fmt.Println(piscine_go.IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(piscine_go.IsCapitalized(""))
}
