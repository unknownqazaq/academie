package _1

func ByeByeFirst(strings []string) []string {

	arr := []string{}

	for i := 1; i < len(strings); i++ {
		arr = append(arr, strings[i])
	}
	return arr
}
