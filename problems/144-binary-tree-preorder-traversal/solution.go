func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	answer := []int{}
	stacks := []*TreeNode{root}

	for len(stacks) > 0 {
		node := stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]

		answer = append(answer, node.Val)

		if node.Right != nil {
			stacks = append(stacks, node.Right)
		}
		if node.Left != nil {
			stacks = append(stacks, node.Left)
		}
	}

	return answer
}

