func luckyNumbers(matrix [][]int) []int {
	minMap := map[int]int{}
	result := []int{}

	for _, row := range matrix {
		min := row[0]

		for _, column := range row {
			if column < min {
				min = column
			}
		}

		minMap[min] = 1
	}

	for i := 0; i < len(matrix[0]); i++ {
		max := 0
		for j := 0; j < len(matrix); j++ {
			n := matrix[j][i]

			if n > max {
				max = n
			}
		}

		if _, ok := minMap[max]; ok {

			result = append(result, max)
		}
	}

	return result
}
