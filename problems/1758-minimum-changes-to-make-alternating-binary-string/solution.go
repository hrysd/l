func minOperations(s string) int {
	count := 0

	for i, c := range s {
		b := int(c - '0')
		if b != i%2 {
			count += 1
		}
	}

	if count > len(s)-count {
		return len(s) - count
	} else {
		return count
	}
}
