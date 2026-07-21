package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	printInt(points.x)
	printStr(", y = ")
	printInt(points.y)
	z01.PrintRune('\n')
}

func printInt(n int) {
	a := '0'
	b := '0'

	for i := 0; i < n/10; i++ {
		a++
	}
	for i := 0; i < n%10; i++ {
		b++
	}
	z01.PrintRune(a)
	z01.PrintRune(b)
}
