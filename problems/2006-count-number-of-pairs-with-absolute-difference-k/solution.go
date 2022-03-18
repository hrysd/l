func countKDifference(nums []int, k int) int {
	result := 0

	m := map[int]int{}

	for _, n := range nums {
		plus, ok := m[n+k]

		if ok {
			result += plus
		}

		minus, ok2 := m[n-k]

		if ok2 {
			result += minus
		}

		m[n] += 1
	}

	return result
}
