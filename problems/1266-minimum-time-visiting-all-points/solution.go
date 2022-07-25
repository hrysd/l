func minTimeToVisitAllPoints(points [][]int) int {
	count := 0

	for i := 1; i < len(points); i++ {
		x1 := points[i-1][0]
		x2 := points[i][0]

		x := int(math.Abs(float64(x2 - x1)))

		y1 := points[i-1][1]
		y2 := points[i][1]

		y := int(math.Abs(float64(y2 - y1)))

		if x > y {
			count += x
		} else {
			count += y
		}
	}

	return count
}
