func majorityElement(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		count, ok := m[v]

		if !ok {
			m[v] = 1
		} else {
			m[v] = count + 1
		}
	}

	var result int
	var count int
	for k, v := range m {
		if v > count {
			count = v
			result = k
		}
	}

	return result
}
