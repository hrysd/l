func countOperations(num1 int, num2 int) int {
	count := 0

	for num1*num2 > 0 {
		if num1 > num2 {
			num1, num2 = num2, num1
		}

		count += num2 / num1
		num2 = num2 % num1
	}

	return count
}
