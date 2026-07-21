package piscine

func Rot14(s string) string {
	result := []rune{}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result = append(result, 'a'+(r-'a'+14)%26)
		} else if r >= 'A' && r <= 'Z' {
			result = append(result, 'A'+(r-'A'+14)%26)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
