func tribonacci(n int) int {
	memo := map[int]int{}

	return tribonacciWithMemo(n, memo)
}

func tribonacciWithMemo(n int, memo map[int]int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	v, ok := memo[n]

	if ok {
		return v
	}

	memo[n] = tribonacciWithMemo(n-3, memo) + tribonacciWithMemo(n-2, memo) + tribonacciWithMemo(n-1, memo)

	return memo[n]
}


