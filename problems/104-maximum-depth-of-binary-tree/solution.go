func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	if leftMax > rightMax {
		max = leftMax
	} else {
		max = rightMax
	}

	return 1 + max // self + depth
}
