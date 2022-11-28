func makeFancyString(s string) string {
	result := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if i > 0 && i < len(s)-1 {
			bc := s[i-1]
			ac := s[i+1]

			if bc == c && c == ac {
				continue
			}
		}

		result = append(result, c)
	}

	return string(result)
}
