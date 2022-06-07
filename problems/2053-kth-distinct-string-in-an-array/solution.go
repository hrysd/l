func kthDistinct(arr []string, k int) string {
	keys := []string{}
	m := map[string]int{}

	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
			keys = append(keys, v)
		}
	}

	index := 1
	for _, key := range keys {
		if m[key] == 1 {
			if index == k {
				return key
			} else {
				index++
			}
		}
	}

	return ""
}
