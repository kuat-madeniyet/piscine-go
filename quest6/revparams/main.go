package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	for i := range os.Args {
		count = i
	}
	for i := count; i > 0; i-- {
		for _, s := range os.Args[i] {
			z01.PrintRune(s)
		}
		z01.PrintRune(10)
	}
}
