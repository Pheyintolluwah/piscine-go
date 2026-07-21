package piscine

func BasicJoin(elems []string) string {
	var result string

	for _, str := range elems {
		result = result + str
	}

	return result
}
