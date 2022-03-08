func numberOfSteps(num int) int {
	count := 0

	for num != 0 {
		if num%2 > 0 {
			num -= 1

		} else {
			num = num / 2

		}

		count++
	}

	return count
}
