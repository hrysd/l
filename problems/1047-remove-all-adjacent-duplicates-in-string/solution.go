func removeDuplicates(s string) string {
	stack := []rune{}
	for _, char := range s {
		l := len(stack)

		if l > 0 && stack[l-1] == char {
			stack = stack[:l-1]
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}
