func checkIfPangram(sentence string) bool {
	chars := map[rune]int{}

	for _, s := range sentence {
		chars[s] = 1
	}

	return len(chars) == 26
}
