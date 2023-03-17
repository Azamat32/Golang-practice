package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	for root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}
