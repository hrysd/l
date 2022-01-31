func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	result := false

	for _, n := range nums {
		v, ok := m[n]
		var count int

		if ok {
			count = v + 1
		} else {
			count = 1
		}

		m[n] = count

		if count > 1 {
			result = true
			break
		}

	}

	return result
}
