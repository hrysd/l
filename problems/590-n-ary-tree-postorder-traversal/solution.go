func postorder(root *Node) []int {
	result := []int{}
	stack := []*Node{root}

	if root == nil {
		return result
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{node.Val}, result...)

		for _, c := range node.Children {
			stack = append(stack, c)
		}
	}

	return result
}
