func getLucky(s string, k int) int {
	answer := 0

	for _, c := range s {
		n := int((c - 'a') + 1)
		answer += (n / 10) + (n % 10)
	}

	for k > 1 && answer > 9 {
		answer = sum(answer)
		k -= 1
	}

	return answer
}

func sum(n int) int {
	answer := 0

	for n > 0 {
		answer += n % 10
		n /= 10
	}

	return answer
}
