func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	words := strings.Split(sentence, " ")

	for i := 1; i < len(words); i++ {
		before := words[i-1][len(words[i-1])-1]
		after := words[i][0]

		if before != after {
			return false
		}
	}

	return true
}
