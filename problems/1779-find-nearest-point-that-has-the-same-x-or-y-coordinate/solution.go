func nearestValidPoint(x int, y int, points [][]int) int {
	min := math.MaxInt
	index := -1

	for i, p := range points {
		if p[0] == x || p[1] == y {
			distance := int(math.Abs(float64(x)-float64(p[0])) + math.Abs(float64(y)-float64(p[1])))

			if distance < min {
				min = distance
				index = i
			}
		}
	}

	return index
}
