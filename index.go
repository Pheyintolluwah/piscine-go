package piscine

func Index(s string, toFind string) int {
	lenS := 0
	lenF := 0

	for range s {
		lenS++
	}
	for range toFind {
		lenF++
	}

	if lenF == 0 {
		return 0
	}

	for i := 0; i <= lenS-lenF; i++ {
		match := true
		for j := 0; j < lenF; j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
