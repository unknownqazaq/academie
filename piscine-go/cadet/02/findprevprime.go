package piscine_go

func FindPrevPrime(nb int) int {
	for i := nb; i > 1; i-- {
		if isprime(i) {
			return i
		}

	}
	return 0

}

func isprime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
