func isUnivalTree(root *TreeNode) bool {
	return check(root, root.Val)
}

func check(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val != val {
		return false
	} else {
		return check(root.Left, val) && check(root.Right, val)
	}
}
