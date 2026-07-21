package piscine

func Split(s, sep string) []string {
	var result []string
	n := len(sep)

	if n == 0 {
		return []string{s}
	}

	start := 0
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == sep {
			result = append(result, s[start:i])
			start = i + n
			i = start - 1
		}
	}

	result = append(result, s[start:])

	return result
}
