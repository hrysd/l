func countTriples(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		j := i

		for i*i+j*j <= n*n {
			sC := i*i + j*j
			c := int(math.Sqrt(float64(sC)))

			if c*c == sC {
				result += 2
			}

			j++
		}
	}

	return result
}
