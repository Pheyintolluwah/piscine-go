package piscine

func UltimateDivMod(a *int, b *int) {
	valA := *a
	valB := *b

	*a = valA / valB
	*b = valA % valB
}
