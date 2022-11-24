func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	count := 0

	for i := 0; i <= len(s)-k; i++ {
		pair, _ := strconv.Atoi(string(s[i : i+k]))

		if pair > 0 && num%pair == 0 {
			count += 1
		}
	}

	return count
}
