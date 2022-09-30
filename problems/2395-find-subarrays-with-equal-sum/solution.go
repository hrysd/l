func findSubarrays(nums []int) bool {
	m := map[int]bool{}

	for i := 1; i < len(nums); i++ {
		n := nums[i] + nums[i-1]

		if _, ok := m[n]; ok {
			return true
		} else {
			m[n] = true
		}
	}

	return false
}
