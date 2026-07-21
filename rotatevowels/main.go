package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	var allRunes []rune
	for i, arg := range args {
		for _, r := range arg {
			allRunes = append(allRunes, r)
		}
		if i < len(args)-1 {
			allRunes = append(allRunes, ' ')
		}
	}

	var vowelIndices []int
	for i, r := range allRunes {
		if isVowel(r) {
			vowelIndices = append(vowelIndices, i)
		}
	}

	for i, j := 0, len(vowelIndices)-1; i < j; i, j = i+1, j-1 {
		idx1 := vowelIndices[i]
		idx2 := vowelIndices[j]
		allRunes[idx1], allRunes[idx2] = allRunes[idx2], allRunes[idx1]
	}

	for _, r := range allRunes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
