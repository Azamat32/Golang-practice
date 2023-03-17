package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	for root.Right == nil {
		return root
	}
	return BTreeMax(root.Right)
}
