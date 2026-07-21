package piscine

func Compact(ptr *[]string) int {
	var nonZero []string
	for _, val := range *ptr {
		if val != "" {
			nonZero = append(nonZero, val)
		}
	}
	*ptr = nonZero
	return len(nonZero)
}
