func timeRequiredToBuy(tickets []int, k int) int {
	seconds := 0

	for tickets[k] > 0 {
		for i, t := range tickets {
			if t > 0 {
				tickets[i] = t - 1
				seconds += 1
			}

			if i == k && tickets[k] == 0 {
				return seconds
			}
		}
	}

	return seconds
}
