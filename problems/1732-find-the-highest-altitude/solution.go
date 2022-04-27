func largestAltitude(gain []int) int {
	current := 0
	max := 0

	for _, v := range gain {
		current = current + v

		if current > max {
			max = current
		}
	}

	return max
}
