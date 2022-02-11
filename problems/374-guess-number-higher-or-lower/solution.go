func guessNumber(n int) int {
	start, end := 0, n

	for start < end {
		number := start + (end-start)/2
		result := guess(number)

		if result == 0 {
			return number
		}

		if result == 1 {
			start = number + 1
		} else {
			end = number - 1
		}
	}

	return start
}
