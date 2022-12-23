func findMaxK(nums []int) int {
	m := map[int]int{}
	max := -1

	for _, n := range nums {
		_, ok := m[n*-1]

		if ok {
			abs := int(math.Abs(float64(n)))
			if abs > max {
				max = abs
			}
		} else {
			m[n] = n
		}
	}

	return max
}
