func isPathCrossing(path string) bool {
	x := 0
	y := 0

	m := map[string]bool{}
	m["0,0"] = true
	for _, p := range path {
		if p == 'N' {
			y += 1
		} else if p == 'S' {
			y -= 1
		} else if p == 'E' {
			x += 1
		} else {
			x -= 1
		}
		point := fmt.Sprintf("%d,%d", x, y)

		if _, ok := m[point]; ok {
			return true
		} else {
			m[point] = true
		}
	}

	return false
}
