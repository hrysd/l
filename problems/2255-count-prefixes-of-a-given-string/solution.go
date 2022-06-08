func countPrefixes(words []string, s string) int {
	count := 0

	for _, w := range words {
		wordLength := len(w)
		length := len(s)
		var l int
		if wordLength > length {
			l = length
		} else {
			l = wordLength
		}

		if w == s[0:l] {
			count++
		}
	}

	return count
}
