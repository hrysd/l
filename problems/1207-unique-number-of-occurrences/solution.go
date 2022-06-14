func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}

	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	c := map[int]int{}

	for _, v := range m {
		if _, ok := c[v]; !ok {
			c[v] = 1
		} else {
			return false
		}
	}

	return true
}
