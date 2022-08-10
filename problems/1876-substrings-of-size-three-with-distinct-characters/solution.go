func countGoodSubstrings(s string) int {
	result := 0

	for i := 3; i <= len(s); i++ {
		a := s[i-3]
		b := s[i-2]
		c := s[i-1]

		if a != b && b != c && c != a {
			result += 1
		}
	}

	return result
}
