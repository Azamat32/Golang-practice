package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	node := l.Head

	for node != nil {

		if comp(ref, node.Data) == true {
			return &node.Data
		}
		node = node.Next
	}
	return &node.Data
}
