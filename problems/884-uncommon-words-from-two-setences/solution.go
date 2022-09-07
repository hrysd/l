func uncommonFromSentences(s1 string, s2 string) []string {
	m := map[string]int{}

	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")

	for _, w := range words1 {
		if _, ok := m[w]; ok {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}
	for _, w := range words2 {
		if _, ok := m[w]; ok {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}

	answer := []string{}

	for w, c := range m {
		if c == 1 {
			answer = append(answer, w)
		}
	}

	return answer
}
