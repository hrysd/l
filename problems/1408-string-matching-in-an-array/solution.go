func stringMatching(words []string) []string {
	m := map[string]bool{}

	for _, w := range words {
		for _, w2 := range words {
			if w == w2 {
				continue
			}
			if strings.Index(w, w2) != -1 {
				m[w2] = true
			}
		}
	}
	answer := []string{}
	for k := range m {
		answer = append(answer, k)
	}

	return answer
}
