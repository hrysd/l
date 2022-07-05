func findDifference(nums1 []int, nums2 []int) [][]int {
	m := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for _, n := range nums1 {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		}
	}

	for _, n := range nums2 {
		if _, ok := m2[n]; !ok {
			m2[n] = struct{}{}
		}
	}

	result1 := []int{}
	for n, _ := range m {
		if _, ok := m2[n]; !ok {
			result1 = append(result1, n)
		}
	}

	result2 := []int{}
	for n, _ := range m2 {
		if _, ok := m[n]; !ok {
			result2 = append(result2, n)
		}
	}

	return [][]int{result1, result2}
}
