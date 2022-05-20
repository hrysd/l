func sumOfUnique(nums []int) int {
	m := map[int]int{}
	result := 0

	for _, v := range nums {
		_, exist := m[v]

		if !exist {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	for k, v := range m {
		if v == 1 {
			result += k
		}
	}

	return result
}
