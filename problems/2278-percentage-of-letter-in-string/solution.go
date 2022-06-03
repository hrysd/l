func percentageLetter(s string, letter byte) int {
	length := len(s)
	count := 0

	for i, _ := range s {
		if s[i] == letter {
			count++
		}
	}

	return int(math.Floor((float64(count) / float64(length)) * 100))
}
