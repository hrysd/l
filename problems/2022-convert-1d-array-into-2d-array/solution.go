func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)

	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < len(original); i++ {
		result[i/n][i%n] = original[i]
	}

	return result
}
