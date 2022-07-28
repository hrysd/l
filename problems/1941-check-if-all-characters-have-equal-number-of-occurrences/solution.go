func areOccurrencesEqual(s string) bool {
	m := map[rune]int{}

	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	occurrence, _ := m[rune(s[0])]

	for _, v := range m {
		if occurrence != v {
			return false
		}
	}

	return true
}
