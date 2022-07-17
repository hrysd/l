func countCharacters(words []string, chars string) int {
	m := map[rune]int{}
	result := 0

	for _, c := range chars {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	for _, word := range words {
		m2 := map[rune]int{}
		canForm := true

		for _, c := range word {
			if _, ok := m2[c]; ok {
				m2[c] += 1
			} else {
				m2[c] = 1
			}
		}

		for k, _ := range m2 {
			if m[k]-m2[k] < 0 {
				canForm = false
				break
			}
		}

		if canForm {
			result += len(word)
		}
	}

	return result
}
