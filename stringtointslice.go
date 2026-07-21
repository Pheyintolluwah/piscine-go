package piscine

func StringToIntSlice(s string) []int {
	if s == "" {
		return nil
	}

	var result []int
	for _, r := range s {
		result = append(result, int(r))
	}
	return result
}
