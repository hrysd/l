func countConsistentStrings(allowed string, words []string) int {
	count := 0
	m := map[rune]int{}

	for _, c := range allowed {
		m[c] = 1
	}

	for _, word := range words {
		same := true

		for _, char := range word {
			_, ok := m[char]
			if !ok {
				same = false
				break
			}
		}

		if same {
			count++
		}
	}

	return count
}
