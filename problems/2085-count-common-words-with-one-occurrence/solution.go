func countWords(words1 []string, words2 []string) int {
	m1 := map[string]int{}
	m2 := map[string]int{}

	for _, w := range words1 {
		if _, ok := m1[w]; ok {
			m1[w]++
		} else {
			m1[w] = 1
		}
	}

	for _, w := range words2 {
		if _, ok := m2[w]; ok {
			m2[w]++
		} else {
			m2[w] = 1
		}
	}

	count := 0
	for _, w := range words1 {
		if m1[w] == 1 && m2[w] == 1 {
			count++
		}
	}
	return count
}
