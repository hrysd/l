func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	return t(root, result)
}

func t(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	result = t(root.Left, result)
	result = t(root.Right, result)

	return append(result, root.Val)
}
