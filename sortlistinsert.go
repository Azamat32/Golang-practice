package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}

	if l == nil {
		l = n
		return l
	}

	iterator := l

	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	ListSort(l)
	return l
}
