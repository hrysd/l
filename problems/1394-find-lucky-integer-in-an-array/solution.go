func findLucky(arr []int) int {
	m := map[int]int{}

	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	answer := -1
	for k, v := range m {
		if k == v {
			if k > answer {
				answer = k
			}
		}
	}

	return answer
}
