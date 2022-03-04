func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 0 {
		number := n % 10

		product *= number
		sum += number

		n /= 10
	}

	return product - sum
}
