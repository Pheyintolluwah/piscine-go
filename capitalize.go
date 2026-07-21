package piscine

func Capitalize(s string) string {
	res := []rune(s)
	newWord := true

	for i := 0; i < len(res); i++ {
		if isAlphanumeric(res[i]) {
			if newWord {
				if res[i] >= 'a' && res[i] <= 'z' {
					res[i] -= 32
				}
				newWord = false
			} else {
				if res[i] >= 'A' && res[i] <= 'Z' {
					res[i] += 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(res)
}

func isAlphanumeric(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}
