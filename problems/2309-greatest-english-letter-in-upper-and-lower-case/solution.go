func greatestLetter(s string) string {
	m := map[rune]int{}

	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 1
		}
	}

	result := ""

	for i := 'Z'; i >= 'A'; i-- {
		_, ok := m[i]
		_, ok2 := m['a'+(i-'A')]

		if ok && ok2 {
			result += string(i)
			break
		}
	}

	return result
}
