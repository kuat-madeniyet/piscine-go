package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, s := range os.Args[0] {
		z01.PrintRune(s)
	}
	z01.PrintRune(10)
}
