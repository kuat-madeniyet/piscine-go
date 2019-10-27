package piscine

func IterativeFactorial(nb int) int {
	sum0 := 1

	if nb < 0 || nb >= 2147483647 {
		return 0
	}

	for i := 1; i < nb+1; i++ {
		sum0 = sum0 * i
	}
	return sum0
}
