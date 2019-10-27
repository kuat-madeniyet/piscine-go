package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	array := []string(os.Args[1:])
	count := 1
	for i := range array {
		count = i
	}
	count = count + 1
	quicksort(array)
	for i := 0; i < count; i++ {
		for _, s := range array[i] {
			z01.PrintRune(s)
		}
		z01.PrintRune(10)
	}
}

func quicksort(table []string) {
	beg := 0
	count := 1
	for i := range table {
		count = i
	}
	count = count + 1

	end := count - 1

	Srot2(table, beg, end)
}
func Srot2(table []string, beg int, end int) {

	if beg < end {
		lockPivo := move(table, beg, end)
		Srot2(table, beg, lockPivo-1)
		Srot2(table, lockPivo+1, end)

	}
}

func move(table []string, beg int, end int) int {
	pivo := table[end]
	i := beg - 1

	for j := beg; j < end; j++ {
		if table[j] <= pivo {
			i++
			aux := table[j]
			table[j] = table[i]
			table[i] = aux
		}
	}

	aux := table[end]
	table[end] = table[i+1]
	table[i+1] = aux

	return i + 1
}
