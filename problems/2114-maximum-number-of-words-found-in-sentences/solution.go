func mostWordsFound(sentences []string) int {
	var maximum int

	for _, v := range sentences {
		count := len(strings.Split(v, " "))

		if count > maximum {
			maximum = count
		}
	}

	return maximum
}

