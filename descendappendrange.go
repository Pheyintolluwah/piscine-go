package piscine

func DescendAppendRange(start, end int) []int {
	result := []int{} // Properly initialized as an empty slice

	for i := start; i > end; i-- {
		result = append(result, i)
	}

	return result
}
