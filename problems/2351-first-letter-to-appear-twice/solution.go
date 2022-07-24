func repeatedCharacter(s string) byte {
	m := map[rune]int{}

	for _, c := range s {
		if count, ok := m[c]; ok {
			if count == 1 {
				return byte(c)
			}

			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	return byte('a')
}
