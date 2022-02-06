import "strings"

func wordPattern(pattern string, s string) bool {
	chars := strings.Split(pattern, "")
	words := strings.Fields(s)

	memo := map[string]string{}
	marked := map[string]string{}

	if len(chars) != len(words) {
		return false
	}

	for i, word := range words {
		w, ok := memo[chars[i]]

		if ok {
			marked[word] = chars[i]

			if w != word {
				return false
			}
		} else {
			memo[chars[i]] = word
			c, ok := marked[word]

			if ok {
				if c != chars[i] {
					return false
				}
			} else {
				marked[word] = chars[i]
			}

		}
	}

	return true
}
