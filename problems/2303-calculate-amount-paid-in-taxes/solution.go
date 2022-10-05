func calculateTax(brackets [][]int, income int) float64 {
	var answer float64

	for i, b := range brackets {
		if income == 0 {
			break
		}
		var d int
		if i > 0 {
			d = b[0] - brackets[i-1][0]
		} else {
			d = b[0]
		}

		if d <= income {
			income -= d
		} else {
			d = income
			income = 0
		}

		answer += float64(d) * (float64(b[1]) / float64(100))
	}

	return answer
}
