func removeOuterParentheses(s string) string {
	result := []string{}
	count := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count += 1
		} else {
			count -= 1
		}

		if count == 0 {
			result = append(result, s[start+1:i])
			start = i + 1
		}
	}

	return strings.Join(result, "")
}
