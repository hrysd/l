func capitalizeTitle(title string) string {
	s := strings.Split(title, " ")

	for i := 0; i < len(s); i++ {
		word := s[i]

		if len(word) > 2 {
			s[i] = strings.Title(strings.ToLower(word))
		} else {
			s[i] = strings.ToLower(word)
		}
	}

	return strings.Join(s, " ")
}
