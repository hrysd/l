func average(salary []int) float64 {
	max := math.MinInt
	min := math.MaxInt
	sum := 0

	for _, s := range salary {
		if s > max {
			max = s
		}

		if s < min {
			min = s
		}

		sum += s
	}

	return float64(sum-(max+min)) / float64(len(salary)-2)
}
