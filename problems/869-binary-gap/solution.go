func binaryGap(n int) int {
	binary := fmt.Sprintf(strconv.FormatInt(int64(n), 2))

	a := 0
	b := 0
	answer := math.MinInt

	for i, c := range binary {
		if c == '1' {
			a = b
			b = i

			if answer < b-a {
				answer = b - a
			}
		}
	}
	return answer
}
