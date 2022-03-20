func getDecimalValue(head *ListNode) int {
	result := 0

	for head != nil {
		v := head.Val
		result = result<<1 | v
		head = head.Next
	}

	return result
}
