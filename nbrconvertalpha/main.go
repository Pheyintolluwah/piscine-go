package main

import (
	"os"

	"github.com/01-edu/z01"
)

func basicAtoi(s string) (int, bool) {
	res := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		res = res*10 + int(r-'0')
	}
	return res, true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := false
	startIdx := 0

	if args[0] == "--upper" {
		upper = true
		startIdx = 1
	}

	for i := startIdx; i < len(args); i++ {
		num, isValid := basicAtoi(args[i])

		if isValid && num >= 1 && num <= 26 {
			if upper {
				z01.PrintRune(rune(64 + num))
			} else {
				z01.PrintRune(rune(96 + num))
			}
		} else {
			z01.PrintRune(' ')
		}
	}

	if len(args) > 0 {
		z01.PrintRune('\n')
	}
}
