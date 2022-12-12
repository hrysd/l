func commonFactors(a int, b int) int {
	var n int

	if a > b {
		n = b
	} else {
		n = a
	}

	count := 0
	for i := 1; i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			count += 1
		}
	}

	return count
}
