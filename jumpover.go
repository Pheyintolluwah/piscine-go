package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	var result []rune
	runes := []rune(str)

	for i := 2; i < len(runes); i += 3 {
		result = append(result, runes[i])
	}

	return string(result) + "\n"
}
