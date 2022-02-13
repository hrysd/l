func firstUniqChar(s string) int {
	split := strings.Split(s, "")
	chars := map[string]int{}
	keys := []string{}

	for _, v := range split {
		keys = append(keys, v)

		_, ok := chars[v]

		if ok {
			chars[v] += 1
		} else {
			chars[v] = 1
		}
	}

	for i, key := range keys {
		count, _ := chars[key]

		if count == 1 {
			return i
		}
	}

	return -1
}
