package piscine

func AtoiBase(s string, base string) int {
	baseLen := len(base)

	if baseLen < 2 {
		return 0
	}

	baseMap := make(map[rune]int)
	for i, char := range base {
		if char == '+' || char == '-' {
			return 0
		}
		if _, exists := baseMap[char]; exists {
			return 0
		}
		baseMap[char] = i
	}

	result := 0
	for _, char := range s {
		val, exists := baseMap[char]
		if !exists {
			return 0
		}
		result = result*baseLen + val
	}

	return result
}
