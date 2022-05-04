func halvesAreAlike(s string) bool {
	length := len(s)
	firstCount := 0
	secondCount := 0

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if isVowels(s[i]) {
			firstCount++
		}

		if isVowels(s[j]) {
			secondCount++
		}
	}

	return firstCount == secondCount
}

func isVowels(b byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, v := range vowels {
		if v == b {
			return true
		}
	}

	return false
}
