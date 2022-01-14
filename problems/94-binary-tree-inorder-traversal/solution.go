func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int

	cursor := root

	for cursor != nil || len(stack) > 0 {
		for cursor != nil {
			stack = append(stack, cursor)
			cursor = cursor.Left
		}

		cursor = stack[len(stack)-1]
		result = append(result, cursor.Val)
		stack = stack[:len(stack)-1]
		cursor = cursor.Right

	}

	return result
}
