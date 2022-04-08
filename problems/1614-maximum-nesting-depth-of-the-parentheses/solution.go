func maxDepth(s string) int {
	depth := 0
	current := 0

	for _, v := range s {
		if v == '(' {
			current++
		} else if v == ')' {
			current--
		}

		if current > depth {
			depth = current
		}
	}

	return depth
}
