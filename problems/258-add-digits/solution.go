func addDigits(num int) int {
	var result int

	for true {
		a := num % 10
		b := num / 10

		result = a + b

		if result > 9 {
			num = result
		} else {
			break
		}
	}

	return result
}
