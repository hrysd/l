func intersection(nums [][]int) []int {
	m := map[int]int{}

	for _, n := range nums {
		for _, v := range n {
			if _, ok := m[v]; !ok {
				m[v] = 1
			} else {
				m[v] += 1
			}
		}
	}
	result := []int{}
	for k, v := range m {
		if v == len(nums) {
			result = append(result, k)
		}
	}

	sort.Ints(result)

	return result
}
