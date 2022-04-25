func firstPlindrome(words []string) string {
	for _, word := range words {
		start := 0
		end := len(word) - 1

		for start <= end {
			if word[start] != word[end] {
				break
			}

			start += 1
			end -=1
		}

		if start >= end {
			return word
		}
	}

	return ""
}a
