func findPoisonedDuration(timeSeries []int, duration int) int {
	seconds := 0

	for i := 1; i < len(timeSeries); i++ {
		a := timeSeries[i] - timeSeries[i-1]

		if a > duration {
			seconds += duration
		} else {
			seconds += a
		}
	}

	return seconds + duration
}
