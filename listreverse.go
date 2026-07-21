package piscine

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head

	newTail := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
	l.Tail = newTail
}
