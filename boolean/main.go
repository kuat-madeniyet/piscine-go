package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)
	for _, letter := range arrayStr {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg(os.Args[1:])) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func lengthOfArg(arg []string) int {
	count := 0
	for range arg {
		count++
	}
	return count
}
