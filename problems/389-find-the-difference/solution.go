func findTheDifference(s string, t string) byte {
	var b byte

	for i := 0; i < len(s); i++ {
		b = b ^ s[i]
	}

	for i := 0; i < len(t); i++ {
		b = b ^ t[i]
	}

	return b
}
