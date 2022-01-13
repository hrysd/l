func deleteDuplicates(head *ListNode) *ListNode {
	var pick func(h *ListNode)
	pick = func(h *ListNode) {
		if h == nil || h.Next == nil {
			return
		}

		if h.Val == h.Next.Val {
			h.Next = h.Next.Next
			pick(h)
		} else {
			pick(h.Next)
		}
	}

	pick(head)

	return head
}
