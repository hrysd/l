func numberOfMatches(n int) int {
	matches := 0

	for n != 1 {
		m := n / 2
		teams := m

		if n%2 != 0 {
			teams += 1
		}

		matches += m
		n = teams
	}

	return matches
}
