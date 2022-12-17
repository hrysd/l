func preorder(root *Node) []int {
	answer := []int{}
	stack := []*Node{root}

	if root == nil {
		return answer
	}

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		answer = append(answer, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append([]*Node{node.Children[i]}, stack...)
		}
	}

	return answer
}
