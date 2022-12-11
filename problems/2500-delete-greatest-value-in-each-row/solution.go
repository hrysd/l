func deleteGreatestValue(grid [][]int) int {
	for i, g := range grid {
		sort.Sort(sort.Reverse(sort.IntSlice(g)))
		grid[i] = g
	}

	answer := 0
	for i := 0; i < len(grid[0]); i++ {
		max := math.MinInt

		for j := 0; j < len(grid); j++ {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}

		answer += max
	}

	return answer
}
