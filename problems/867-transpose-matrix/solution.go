func transpose(matrix [][]int) [][]int {
	rowCount := len(matrix[0])
	columnCount := len(matrix)

	rows := make([][]int, rowCount)

	for i, _ := range rows {
		rows[i] = make([]int, columnCount)
	}

	for i, row := range matrix {
		for j, col := range row {
			rows[j][i] = col
		}
	}

	return rows
}
