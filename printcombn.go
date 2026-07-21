package piscine

import "github.com/01-edu/z01"

func printNumber(arr []int) {
	for _, v := range arr {
		z01.PrintRune(rune(v) + '0')
	}
}

func generate(arr []int, index int, start int, n int) {
	if index == n {
		printNumber(arr)

		if arr[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		arr[index] = i
		generate(arr, index+1, i+1, n)
	}
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	arr := make([]int, n)
	generate(arr, 0, 0, n)
	z01.PrintRune('\n')
}
