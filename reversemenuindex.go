package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversed := make([]string, length)
	for i, item := range menu {
		reversed[length-1-i] = item
	}
	return reversed
}
