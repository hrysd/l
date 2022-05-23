func countNegatives(grid [][]int) int {
	count := 0

	for _, i := range grid {
		count += binarySearch(i)
	}

	return count
}

func binarySearch(s []int) int {
	i := 0
	position := len(s)

	for i < position {
		middle := i + (position-i)/2

		if s[middle] < 0 {
			position = middle
		} else {
			i = middle + 1
		}
	}

	return len(s) - position
}
