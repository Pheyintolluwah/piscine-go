package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	var strToProcess string
	var insertStr string
	orderFlag := false

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			strToProcess = arg
		}
	}

	resultStr := strToProcess + insertStr

	if orderFlag {
		runes := []rune(resultStr)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		resultStr = string(runes)
	}

	for _, r := range resultStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
