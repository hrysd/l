func checkAlmostEquivalent(word1 string, word2 string) bool {
	m := map[rune]int{}

	for _, w := range word1 {
		if _, ok := m[w]; ok {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}

	for _, w := range word2 {
		if _, ok := m[w]; ok {
			m[w] -= 1
		} else {
			m[w] = -1
		}
	}

	for _, v := range m {
		c := math.Abs(float64(v))
		if int(c) > 3 {
			return false
		}
	}

	return true
}
