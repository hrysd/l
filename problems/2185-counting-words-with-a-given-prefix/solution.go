func prefixCount(words []string, pref string) int {
	count := 0
	re := regexp.MustCompile(fmt.Sprintf("\\A%s", pref))

	for _, word := range words {
		if re.MatchString(word) {
			count++
		}
	}

	return count
}
