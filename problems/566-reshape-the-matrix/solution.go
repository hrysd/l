func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	answer := [][]int{}

	for i := 0; i < r; i++ {
		answer = append(answer, []int{})
	}

	row := 0

	for _, m := range mat {
		for _, v := range m {
			if len(answer[row]) == c {
				row += 1
			}

			answer[row] = append(answer[row], v)
		}
	}

	return answer
}
