func numIdenticalPairs(nums []int) int {
	m := map[int]int{}
	count := 0

	for i := 0; i < len(nums); i++ {
		c, ok := m[nums[i]]

		if ok {
			count += c + 1
			m[nums[i]] = c + 1
		} else {
			m[nums[i]] = 0
		}
	}

	return count
}
