package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1
	foundDigit := false

	for _, char := range s {
		if char == '-' && !foundDigit {
			sign = -1
		} else if char >= '0' && char <= '9' {
			foundDigit = true

			digit := int(char - '0')
			res = res*10 + digit
		}
	}

	return res * sign
}
