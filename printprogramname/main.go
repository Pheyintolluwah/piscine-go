package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	lastSlash := -1
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			lastSlash = i
			break
		}
	}

	for _, r := range name[lastSlash+1:] {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
