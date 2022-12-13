func pivotInteger(n int) int {
	n = n * (n + 1) / 2
	sqrt := int(math.Sqrt(float64(n)))

	if sqrt*sqrt == n {
		return sqrt
	} else {
		return -1
	}
}
