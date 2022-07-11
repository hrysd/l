func frequencySort(nums []int) []int {
	m := map[int]int{}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] += 1
		} else {
			m[n] = 1
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		if m[nums[i]] == m[nums[j]] {
			return nums[i] > nums[j]
		} else {
			return m[nums[i]] < m[nums[j]]
		}
	})

	return nums
}
