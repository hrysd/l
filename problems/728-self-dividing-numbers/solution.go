func selfDividingNumbers(left int, right int) []int {
	result := []int{}

	for i := left; i <= right; i++ {
		n := i

		for n > 0 {
			digit := n % 10

			if digit == 0 || i%digit != 0 {
				break
			}

			n /= 10

			if n == 0 {
				result = append(result, i)
			}
		}
	}

	return result
}
