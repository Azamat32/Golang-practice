package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		currentLevel(root, i, f)
	}
}

func currentLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		currentLevel(root.Left, level-1, f)
		currentLevel(root.Right, level-1, f)
	}
}
