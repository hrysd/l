func isToeplitzMatrix(matrix [][]int) bool {
	row := len(matrix)
	col := len(matrix[0])

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}
