func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	left := root.Left
	if left != nil && left.Left == nil && left.Right == nil {
		sum += left.Val
	}

	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
