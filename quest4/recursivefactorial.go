package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb >= 2147483647 {
		return 0
	}
	if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
