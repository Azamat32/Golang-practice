package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		n1 = n2
		return ListSort(n1)
	} else if n2 == nil {
		n2 = n1
		return ListSort(n2)
	}
	head := n1
	for head.Next != nil {
		head = head.Next
	}
	head.Next = n2
	return ListSort(n1)
}
