package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := len(args) - 1; i >= 1; i-- {
		arg := args[i]

		for _, r := range arg {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
	}
}
