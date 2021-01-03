var PAIR = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isValid(s string) bool {
	sRune := []rune(s)
	var endParentheses []string

	if len(sRune) == 1 {
		return false
	}

	for i, v := range sRune {
		end, exists := PAIR[string(v)]

		if exists {
			endParentheses = append(endParentheses, end)
			continue
		} else {
			if len(endParentheses) == 0 {
				return false
			}

			last := len(endParentheses) - 1
			expected := endParentheses[last]

			if expected != string(sRune[i]) {
				return false
			}

			endParentheses = append(endParentheses[:last], endParentheses[last+1:]...)
		}
	}

	return len(endParentheses) == 0
}
