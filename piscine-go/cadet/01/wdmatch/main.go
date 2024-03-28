package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		wdmatch(os.Args[1], os.Args[2])
	}

	/*wdmatch("quarante deuxq", "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq ")*/
}

func wdmatch(a, b string) {
	result := ""
	index := 0

	for _, char := range a {
		for i := index; i < len(b); i++ {
			if char == rune(b[i]) {
				result += string(char)
				index = i + 1
				break
			}

		}

	}
	if result == a {
		os.Stdout.WriteString(result + "\n")

	}

}
