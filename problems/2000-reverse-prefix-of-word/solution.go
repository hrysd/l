func reversePrefix(word string, ch byte) string {
	tmp := []byte{}
	isOccurred := false

	for i := 0; i < len(word); i++ {
		w := word[i]
		if isOccurred {
			tmp = append(tmp, w)
		} else {
			tmp = append([]byte{w}, tmp...)
		}

		if w == ch {
			isOccurred = true
		}
	}

	if isOccurred {
		return string(tmp)
	} else {
		return word
	}
}
