package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := 0
	baseFromLen := len(baseFrom)

	for _, char := range nbr {
		index := 0
		for i, bChar := range baseFrom {
			if bChar == char {
				index = i
				break
			}
		}
		decimalValue = decimalValue*baseFromLen + index
	}

	if decimalValue == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	var resultBytes []byte

	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		resultBytes = append([]byte{baseTo[remainder]}, resultBytes...)
		decimalValue /= baseToLen
	}

	return string(resultBytes)
}
