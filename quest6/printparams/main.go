package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := range os.Args {
		if i == 0 {
			continue
		} else {
			for _, s := range os.Args[i] {
				z01.PrintRune(s)
			}
			z01.PrintRune(10)
		}
	}
}
