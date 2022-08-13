func commonChars(words []string) []string {
	result := []string{}

	for i := 'a'; i <= 'z'; i++ {
		count := math.MaxInt

		for _, word := range words {
			cCount := 0

			for _, c := range word {
				if i == c {
					cCount += 1
				}
			}

			if count > cCount {
				count = cCount
			}
		}

		for j := 0; j < count; j++ {
			result = append(result, string(i))
		}
	}

	return result
}
