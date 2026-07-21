package piscine

func SortIntegerTable(table []int) {
	a := len(table)
	for i := 0; i < a-1; i++ {
		for j := 0; j < a-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
