func countPrimeSetBits(left int, right int) int {
	result := 0

	for i := left; i <= right; i++ {
		count := bits.OnesCount(uint(i))

		if isPrime(count) {
			result += 1
		}
	}
	return result
}

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}

	count := 0

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			count += 1
			break
		}
	}

	return count == 0
}
