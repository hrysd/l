func divideArray(nums []int) bool {
	m := map[int]int{}

	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}

	return true
}
