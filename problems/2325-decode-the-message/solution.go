func decodeMessage(key string, message string) string {
	m := map[rune]rune{}
	cursor := 'a'

	for _, c := range key {
		if c == ' ' {
			continue
		}

		if _, ok := m[c]; !ok {
			m[c] = cursor
			cursor += 1
		}
	}

	results := []rune{}

	for _, c := range message {
		if c == ' ' {
			results = append(results, ' ')
		} else {
			results = append(results, m[c])
		}
	}

	return string(results)
}
