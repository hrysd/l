func numberOfLines(widths []int, s string) []int {
	lines := 1
	width := 0

	for _, c := range s {
		width += widths[c-'a']

		if width > 100 {
			lines += 1
			width = widths[c-'a']
		}
	}

	return []int{lines, width}
}
