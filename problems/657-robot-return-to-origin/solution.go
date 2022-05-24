func judgeCircle(moves string) bool {
	vertical := 0
	horizontal := 0

	for _, m := range moves {
		if m == 'U' {
			vertical += 1
		} else if m == 'D' {
			vertical -= 1
		} else if m == 'L' {
			horizontal += 1
		} else if m == 'R' {
			horizontal -= 1
		}
	}

	return vertical == 0 && horizontal == 0
}
