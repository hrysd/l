func countAsterisks(s string) int {
	bar := 0
	count := 0

	for _, c := range s {
		if c == '|' {
			bar += 1
			continue
		}

		if bar%2 == 0 && c == '*' {
			count += 1
		}
	}

	return count
}
