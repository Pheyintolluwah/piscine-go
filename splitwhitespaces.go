package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string

	for i, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}

		if i == len(s)-1 && word != "" {
			result = append(result, word)
		}
	}

	return result
}
