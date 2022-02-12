func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	m := map[string]int{}

	for _, v := range strings.Split(magazine, "") {
		i, ok := m[v]

		if ok {
			m[v] = i + 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range strings.Split(ransomNote, "") {
		i, ok := m[v]

		if i == 0 {
			return false
		}

		if ok {
			m[v] = i - 1
		} else {
			return false
		}
	}

	return true
}
