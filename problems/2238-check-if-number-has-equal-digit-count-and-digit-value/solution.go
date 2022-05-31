func digitCount(num string) bool {
	expected := map[int]int{}

	nums := strings.Split(num, "")

	for _, n := range nums {
		e, _ := strconv.Atoi(n)
		if _, ok := expected[e]; !ok {
			expected[e] = 0
		}
		expected[e] += 1
	}

	for i, n := range nums {
		e, _ := strconv.Atoi(n)
		if expected[i] != e {
			return false
		}
	}

	return true
}
