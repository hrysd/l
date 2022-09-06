func distributeCandies(candyType []int) int {
	candies := len(candyType) / 2
	m := map[int]struct{}{}

	for _, c := range candyType {
		if _, ok := m[c]; !ok {
			m[c] = struct{}{}
		}
	}

	if len(m) < candies {
		return len(m)
	} else {
		return candies
	}
}
