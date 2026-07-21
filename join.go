package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]

	for i := 1; i < len(strs); i++ {
		res = res + sep + strs[i]
	}

	return res
}
