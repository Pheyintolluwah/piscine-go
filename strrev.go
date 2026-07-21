package piscine

func StrRev(s string) string {
	// Convert string to a slice of runes to handle UTF-8 correctly
	runes := []rune(s)

	// Swap elements from the outside in
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the rune slice back into a string
	return string(runes)
}
