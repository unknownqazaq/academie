package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		result := sumascii(os.Args[1])
		fmt.Println(result)

	} else {
		fmt.Println(0)
	}

}

func sumascii(str string) int {
	arrByte := []byte(str)
	var result byte
	for _, char := range arrByte {
		result += char
	}
	return int(result)

}
