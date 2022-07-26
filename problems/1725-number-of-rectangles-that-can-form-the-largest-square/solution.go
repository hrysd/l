func countGoodRectangles(rectangles [][]int) int {
	max := 0
	count := 0

	for _, r := range rectangles {
		l, w := r[0], r[1]
		var min int

		if l > w {
			min = w
		} else {
			min = l
		}

		if min > max {
			max = min
			count = 1
		} else if max == min {
			count++
		}
	}

	return count
}
