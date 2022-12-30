func removeAnagrams(words []string) []string {
	var previous string
	answer := []string{}

	for _, word := range words {
		chars := strings.Split(word, "")
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		sorted := strings.Join(chars, "")

		if previous != sorted {
			previous = sorted
			answer = append(answer, word)
		}
	}

	return answer
}
