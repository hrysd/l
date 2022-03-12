func sortSentence(s string) string {
	sentences := strings.Split(s, " ")
	sorted := make([]string, len(sentences))

	for _, sentence := range sentences {
		index, _ := strconv.Atoi(sentence[len(sentence)-1:])

		sorted[index-1] = sentence[:len(sentence)-1]
	}

	return strings.Join(sorted, " ")
}
