package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return validateBst(root, nil, nil)
}

func validateBst(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}

	val := Atoi(root.Data)

	if (min != nil && val <= *min) || (max != nil && val >= *max) {
		return false
	}

	left := validateBst(root.Left, min, &val)
	right := validateBst(root.Right, &val, max)

	return left && right
}
