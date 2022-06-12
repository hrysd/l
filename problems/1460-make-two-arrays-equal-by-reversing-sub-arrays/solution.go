func canBeEqual(target []int, arr []int) bool {
	m := map[int]int{}

	for i, v := range target {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}

		if _, ok := m[arr[i]]; !ok {
			m[arr[i]] = 0
		}

		m[v] += 1
		m[arr[i]] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
