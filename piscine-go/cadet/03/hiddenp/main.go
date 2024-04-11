package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		args := os.Args[1]
		args2 := os.Args[2]
		res := hiddenp(args, args2)
		os.Stdout.WriteString(res + "\n")
	}

}

func hiddenp(s1, s2 string) string {
	res := ""
	index := 0
	for i := 0; i < len(s1); i++ {
		for j := index; j < len(s2); j++ {
			if s1[i] == s2[j] {
				res += string(s1[i])
				index++
				break // Exit the inner loop once a match is found
			}
		}
	}
	if s1 == res {
		return "1"
	} else {
		return "0"
	}

}
