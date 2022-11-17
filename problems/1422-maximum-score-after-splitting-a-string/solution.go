func maxScore(s string) int {
	answer := 0

	for i := 1; i < len(s); i++ {
		left, right := split(s, i)

		current := strings.Count(left, "0") + strings.Count(right, "1")

		answer = max(current, answer)
	}

	return answer
}

func split(s string, n int) (string, string) {
	r := []rune(s)
	l := len(r)
	left := r[0:n]
	right := r[n:l]

	return string(left), string(right)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}


