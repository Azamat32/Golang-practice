package piscine

func ListLast(l *List) interface{} {
	iterator := l.Head
	if iterator == nil {
		return nil
	}
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	return iterator.Data
}
