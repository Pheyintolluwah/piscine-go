package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := os.Args[1]
	data, _ := os.ReadFile(fileName)

	fmt.Print(string(data))
}
