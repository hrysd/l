func countPoints(rings string) int {
	pair := ""
	pairs := []string{}

	for i, ring := range rings {
		pair = pair + string(ring)

		if i > 0 && (i+1)%2 == 0 {
			pairs = append(pairs, pair)
			pair = ""
		}
	}

	m := map[string]map[string]int{}

	for _, p := range pairs {
		_, ok := m[string(p[1])]

		if !ok {
			m[string(p[1])] = map[string]int{}
		}

		m[string(p[1])][string(p[0])] = 1
	}
	result := 0
	for _, a := range m {
		if len(a) == 3 {
			result++
		}
	}
	return result
}
