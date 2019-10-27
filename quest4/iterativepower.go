package piscine

func IterativePower(nb int, power int) int {
	sum0 := 1
	if power < 0 {
		return 0
	}
	for i := 1; i < power+1; i++ {
		sum0 = nb * sum0
	}
	return sum0
}
