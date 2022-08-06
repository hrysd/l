func numberOfPairs(nums []int) []int {
	m := map[int]int{}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] += 1
		} else {
			m[n] = 1
		}
	}

	pairs := 0
	leftover := 0
	for _, v := range m {

		pairs += v / 2
		leftover += v % 2

	}

	return []int{pairs, leftover}
}
