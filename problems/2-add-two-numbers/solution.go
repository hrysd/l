func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: 0}

	right, left, current := l1, l2, root
	carry := 0

	for right != nil || left != nil {
		var rightVal int
		var leftVal int

		if right != nil {
			rightVal = right.Val
		}

		if left != nil {
			leftVal = left.Val
		}

		val := rightVal + leftVal + carry
		carry = val / 10
		node := &ListNode{Val: val % 10}
		current.Next = node
		current = node

		if right != nil {
			right = right.Next
		}

		if left != nil {
			left = left.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return root.Next
}
