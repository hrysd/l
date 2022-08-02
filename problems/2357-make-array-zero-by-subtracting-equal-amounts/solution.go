func minimumOperations(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		if v == 0 {
			continue
		}
		if _, ok := m[v]; !ok {
			m[v] = 1
		}
	}

	return len(m)
}
