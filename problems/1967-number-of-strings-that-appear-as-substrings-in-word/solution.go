func numOfStrings(patterns []string, word string) int {
	var count int

	for _, pattern := range patterns {
		if strings.Index(word, pattern) != -1 {
			count++
		}
	}

	return count
}
