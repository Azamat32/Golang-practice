package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL

	curr := l.Head

	for curr != nil {
		if curr.Data == data_ref {
			if prev == nil {
				l.Head = curr.Next
			} else {
				prev.Next = curr.Next
			}
		} else {
			prev = curr
		}
		curr = curr.Next

	}
}
