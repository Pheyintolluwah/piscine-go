package main

import (
	"fmt"
	"os"
)

func basicAtoi(s string) int {
	res := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		}
	}
	return res
}

func main() {
	if len(os.Args) < 4 {
		return
	}

	n := basicAtoi(os.Args[2])

	files := os.Args[3:]
	hasError := false

	for i, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", fileName)
			hasError = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", fileName)
		}

		stats, _ := file.Stat()
		size := stats.Size()

		start := size - int64(n)
		if start < 0 {
			start = 0
		}

		buffer := make([]byte, n)
		bytesRead, _ := file.ReadAt(buffer, start)

		os.Stdout.Write(buffer[:bytesRead])

		file.Close()
	}

	if hasError {
		os.Exit(1)
	}
}
