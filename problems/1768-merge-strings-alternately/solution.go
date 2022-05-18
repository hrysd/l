func mergeAlternately(word1 string, word2 string) string {
	var max int
	len1 := len(word1)
	len2 := len(word2)

	if len1 > len2 {
		max = len1
	} else {
		max = len2
	}

	result := []byte{}

	for i := 0; i < max; i++ {
		if i < len1 {
			result = append(result, word1[i])
		}

		if i < len2 {
			result = append(result, word2[i])
		}
	}

	return string(result)
}
