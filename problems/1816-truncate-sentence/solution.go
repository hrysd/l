func truncateSentence(s string, k int) string {
	result := []rune{}
	count := 0

	for _, r := range s {
		if r == ' ' {
			count++

			if count == k {
				break
			}
		}

		result = append(result, r)
	}

	return string(result)
}
