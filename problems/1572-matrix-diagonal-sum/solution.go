func diagonalSum(mat [][]int) int {
	var result int
	length := len(mat[0])
	sencodary := length - 1

	for i, m := range mat {
		result += m[i]

		if sencodary-i != i {
			result += m[sencodary-i]
		}
	}

	return result
}
