func countVowelSubstrings(word string) int {
	length := len(word)
	answer := 0

	for i, _ := range word {
		set := map[byte]struct{}{}

		for j := i; j < length; j++ {
			if !isVowel(word[j]) {
				break
			}

			if _, ok := set[word[j]]; !ok {
				set[word[j]] = struct{}{}
			}

			if len(set) == 5 {
				answer += 1

			}
		}
	}

	return answer
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o'
}
