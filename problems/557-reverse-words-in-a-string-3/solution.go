func reverseWords(s string) string {
	splitted := strings.Split(s, " ")

	result := []rune{}

	for i, word := range splitted {
		wordRune := []rune(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			wordRune[i], wordRune[j] = wordRune[j], wordRune[i]
		}
		if i != 0 {
			result = append(result, []rune(" ")...)
		}
		result = append(result, wordRune...)
	}

	return string(result)
}
