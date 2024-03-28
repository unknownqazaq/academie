package main

/*

inter
Instructions
Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

The display will be followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing.

Usage
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
$


*/

func main() {
	//if len(os.Args) == 3 {
	//	word := inter(os.Args[1], os.Args[2])
	//	fmt.Println(word)
	//}

	inter("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
}

func inter(str1, str2 string) string {
	index := 0
	result := ""

	for _, char := range str1 {
		for i := index; i < len(str2); i++ {
			if char == rune(str2[i]) {
				result += string(char)
				index += 1
				break
			}

		}
	}

	return doubles(result)

}

func doubles(str string) string {
	appeared := make(map[rune]bool)
	result := ""

	for _, char := range str {
		if !appeared[char] {
			result += string(char)
			appeared[char] = true
		}
	}
	return result
}
