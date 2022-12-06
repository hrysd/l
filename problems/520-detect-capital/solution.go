func detectCapitalUse(word string) bool {
	allUpper := strings.ToUpper(word)
	allLower := strings.ToLower(word)
	capitalized := strings.Title(allLower)

	return word == allUpper || word == allLower || word == capitalized
}
