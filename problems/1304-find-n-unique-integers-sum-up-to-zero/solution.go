func sumZero(n int) []int {
	result := make([]int, n)

	for i := 1; i < n; i++ {
		result[i] = i
		result[0] -= i
	}

	return result
}

