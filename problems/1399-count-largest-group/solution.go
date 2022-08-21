func countLargestGroup(n int) int {
	max := 0
	m := map[int]int{}

	for i := 1; i <= n; i++ {
		v := sum(i)

		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}

		if max < m[v] {
			max = m[v]
		}
	}

	count := 0
	for _, v := range m {
		if v == max {
			count += 1
		}
	}

	return count
}

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n%10 + sum(n/10)
	}
}
