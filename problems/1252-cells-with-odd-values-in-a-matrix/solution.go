func oddCells(m int, n int, indices [][]int) int {
	count := 0

	row := make([]int, m)
	col := make([]int, n)

	for _, i := range indices {
		row[i[0]] += 1
		col[i[1]] += 1
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {

			if (row[r]+col[c])%2 != 0 {
				count += 1
			}
		}
	}

	return count
}
