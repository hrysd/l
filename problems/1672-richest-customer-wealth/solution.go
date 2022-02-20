func maximumWealth(accounts [][]int) int {
	var maximum int

	for _, account := range accounts {
		var sum int

		for _, v := range account {
			sum += v
		}

		// maximum = int(math.Max(float64(maximum), float64(sum)))
		if sum > maximum {
			maximum = sum
		}
	}

	return maximum
}
