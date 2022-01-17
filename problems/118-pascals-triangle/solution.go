// https://www.mathsisfun.com/pascals-triangle.html
func generate(numRows int) [][]int {
	var result [][]int

	result = append(result, []int{1})
	for numRow := 1; numRow < numRows; numRow++ {
		row := make([]int, numRow+1)
		row[0] = 1
		columns := len(result[numRow-1]) - 1

		for column := 0; column < columns; column++ {
			row[column+1] = result[numRow-1][column] + result[numRow-1][column+1]
		}
		row[len(row)-1] = 1
		result = append(result, row)
	}

	return result
}
