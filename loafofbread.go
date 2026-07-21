package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(str); i++ {
		if count < 5 && str[i] == ' ' {
			continue
		}

		if count == 5 {

			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}

			if i == len(str)-1 {
				break
			}
			result += " "
			count = 0
			continue
		}

		result += string(str[i])
		count++
	}

	result += "\n"
	return result
}
