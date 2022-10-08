func maxNumberOfBalloons(text string) int {
	m := map[rune]int{}

	for _, c := range text {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	return min(m['b'], min(m['a'], min(m['l']/2, min(m['o']/2, m['n']))))
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
