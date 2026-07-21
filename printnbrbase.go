package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := 0
	for range base {
		baseLen++
	}

	n := nbr
	if n < 0 {
		z01.PrintRune('-')
	} else {
		n = -n
	}

	printRecursive(n, base, baseLen)
}

func printRecursive(n int, base string, baseLen int) {
	if n <= -baseLen {
		printRecursive(n/baseLen, base, baseLen)
	}

	index := -(n % baseLen)
	z01.PrintRune(rune(base[index]))
}

func isValidBase(base string) bool {
	runes := []rune(base)
	if len(runes) < 2 {
		return false
	}

	for i := 0; i < len(runes); i++ {
		if runes[i] == '+' || runes[i] == '-' {
			return false
		}
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}
