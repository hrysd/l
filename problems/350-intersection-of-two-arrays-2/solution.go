func intersect(nums1 []int, nums2 []int) []int {
	var a []int
	var b []int

	if len(nums1) > len(nums2) {
		a = nums2
		b = nums1
	} else {
		a = nums1
		b = nums2
	}
	m := map[int]int{}

	for _, n := range a {
		m[n] += 1
	}

	answer := []int{}
	for _, n := range b {
		if v, ok := m[n]; ok && v > 0 {
			answer = append(answer, n)
			m[n] -= 1
		}
	}
	return answer
}
