func isPrefixOfWord(sentence string, searchWord string) int {
	splitted := strings.Split(sentence, " ")

	for i, s := range splitted {
		if strings.HasPrefix(s, searchWord) {
			return i + 1
		}
	}

	return -1
}
