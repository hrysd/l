func checkXMatrix(grid [][]int) bool {
	lastColumn := len(grid[0]) - 1

	for i, g := range grid {
		for j, n := range g {
			if i == j || i+j == lastColumn {
				if n == 0 {
					return false
				}
			} else {
				if n > 0 {
					return false
				}
			}
		}
	}

	return true
}
