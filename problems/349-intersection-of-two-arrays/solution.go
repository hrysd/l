func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var n []int
	var n2 []int
	if len(nums1) > len(nums2) {
		n = nums2
		n2 = nums1
	} else {
		n = nums1
		n2 = nums2
	}

	for _, v := range n {
		m[v] = true
	}

	var result []int

	for _, v := range n2 {
		_, ok := m[v]

		if ok {
			result = append(result, v)
			delete(m, v)
		}
	}

	return result
}
