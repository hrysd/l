func rangeSumBST(root *TreeNode, low int, high int) int {
	a := inorder(root, low, high)

	return a
}

func inorder(node *TreeNode, low int, high int) int {
	sum := 0

	if node == nil {
		return sum
	}

	sum += inorder(node.Left, low, high)

	if node.Val >= low && node.Val <= high {
		sum += node.Val
	}

	sum += inorder(node.Right, low, high)

	return sum
}
