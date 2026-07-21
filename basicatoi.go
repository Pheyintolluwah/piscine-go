package piscine

func BasicAtoi(s string) int {
	res := 0
	for _, r := range s {
		digit := int(r - '0')
		res = res*10 + digit
	}
	return res
}
