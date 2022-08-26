func digitSum(s string, k int) string {
	for {
		if len(s) <= k {
			break
		}
		tmp := ""
		for i := 0; i < len(s); i += k {
			var w string
			if i+k < len(s) {
				w = s[i:(i + k)]
			} else {
				w = s[i:]
			}

			answer := 0
			for _, c := range w {
				answer += int(c - '0')
			}

			tmp += strconv.Itoa(answer)
		}

		s = tmp
	}

	return s
}
