package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l1.Head == nil {
		l1.Head = l2.Head
		return
	} else if l2 == nil || l2.Head == nil {
		return
	}
	head := l1.Head

	for head.Next != nil {
		head = head.Next
	}
	head.Next = l2.Head
}
