func reformat(s string) string {
	if len(s) == 1 {
		return s
	}

	alpha := []rune{}
	digits := []rune{}

	for _, c := range s {
		if c >= '0' && c <= '9' {
			digits = append(digits, c)
		} else {
			alpha = append(alpha, c)
		}
	}

	if math.Abs(float64(len(digits)-len(alpha))) > 1 {
		return ""
	}

	answer := make([]rune, len(s))
	startByDigit := len(digits) > len(alpha)
	for i, a, d := 0, 0, 0; i < len(s); i++ {
		if startByDigit {
			if d < len(digits) {
				answer[i] = digits[d]
				d += 1
			}
		} else {
			if a < len(alpha) {
				answer[i] = alpha[a]
				a += 1
			}
		}

		startByDigit = !startByDigit
	}

	return string(answer)
}
