func findWords(words []string) []string {
	result := []string{}
	row1 := regexp.MustCompile(`\A[qwertyuiopQWERTYUIOP]*\z`)
	row2 := regexp.MustCompile(`\A[asdfghjklASDFGHJKL]*\z`)
	row3 := regexp.MustCompile(`\A[zxcvbnmZXCVBNM]*\z`)

	for _, w := range words {
		if row1.MatchString(w) || row2.MatchString(w) || row3.MatchString(w) {
			result = append(result, w)
		}
	}

	return result
}
