func oddString(words []string) string {
	m := map[string][]int{}

	for index, word := range words {
		diff := []int{}

		for i := 1; i < len(word); i++ {
			diff = append(diff, int(word[i]-'a')-int(word[i-1]-'a'))
		}

		key := fmt.Sprint(diff)

		if _, ok := m[key]; ok {
			m[key] = append(m[key], index)
		} else {
			m[key] = []int{index}
		}
	}

	for _, v := range m {
		if len(v) == 1 {
			return words[v[0]]
		}
	}

	return "unreachable"
}
