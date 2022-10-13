func maxPower(s string) int {
	answer := 1
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count += 1
			if answer < count {
				answer = count
			}
		} else {

			count = 1
		}
	}

	return answer
}
