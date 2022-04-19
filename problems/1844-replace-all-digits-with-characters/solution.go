func replaceDigits(s string) string {
	result := make([]byte, len(s))

	for i := 0; i < len(s); i = i + 2 {
		char := s[i]
		result[i] = char

		if i+1 < len(s) {
			digit := byte(s[i+1] - '0')
			result[i+1] = char + digit
		}
	}

	return string(result)
}
