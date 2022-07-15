func shiftGrid(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])

	result := make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			newColumn := (c + k) % col
			newRow := (r + (c+k)/col) % row

			result[newRow][newColumn] = grid[r][c]
		}
	}

	return result
}
