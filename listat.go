package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}
*/

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	current := l
	for i := 0; i < pos; i++ {
		if current.Next == nil {
			return nil
		}
		current = current.Next
	}

	return current
}
