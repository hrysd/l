func reverseOnlyLetters(s string) string {
	temp := []byte{}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			temp = append(temp, s[i])
		}
	}

	answer := []byte{}

	for i := 0; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i])) {
			answer = append(answer, s[i])
		} else {
			c := temp[0]
			temp = temp[1:]
			answer = append(answer, c)
		}
	}
	return string(answer)
}
