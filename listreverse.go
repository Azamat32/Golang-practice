package piscine

func ListReverse(l *List) {
	curr := l.Head
	var prev *NodeL
	var next *NodeL

	for curr != nil {
		// Store next
		next = curr.Next
		// Reverse current node's pointer
		curr.Next = prev

		// Move pointers one position ahead.
		prev = curr
		curr = next
	}
	l.Head = prev
}
