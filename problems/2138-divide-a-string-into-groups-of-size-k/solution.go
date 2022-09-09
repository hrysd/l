func divideString(s string, k int, fill byte) []string {
	answer := []string{}

	tmp := []rune{}

	for i, c := range s {
		tmp = append(tmp, c)

		if (i+1)%k == 0 {
			answer = append(answer, string(tmp))
			tmp = []rune{}
		}
	}

	if len(tmp) > 0 {
		c := len(tmp)
		for i := 0; i < k-c; i++ {
			tmp = append(tmp, rune(fill))
		}

		answer = append(answer, string(tmp))
	}

	return answer
}
