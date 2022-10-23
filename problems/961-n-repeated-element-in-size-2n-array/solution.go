func repeatedNTimes(nums []int) int {
	m := map[int]bool{}

	for _, v := range nums {
		_, exist := m[v]

		if exist {
			return v
		} else {
			m[v] = true
		}
	}

	return -1
}
