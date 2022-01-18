func getRow(rowIndex int) []int {
	var result [][]int

	result = append(result, []int{1})
	for i := 1; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0] = 1
		columns := len(result[i-1]) - 1

		for column := 0; column < columns; column++ {
			row[column+1] = result[i-1][column] + result[i-1][column+1]
		}
		row[len(row)-1] = 1
		result = append(result, row)
	}

	return result[rowIndex]
}
