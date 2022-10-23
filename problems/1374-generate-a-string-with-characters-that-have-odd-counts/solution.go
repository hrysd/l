func generateTheString(n int) string {
	isOdd := n%2 == 0

	if isOdd {
		return strings.Repeat("a", n-1) + "b"
	} else {
		return strings.Repeat("a", n)
	}
}
