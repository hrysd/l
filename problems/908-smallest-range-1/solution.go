func smallestRangeI(nums []int, k int) int {
	min := math.MaxInt
	max := math.MinInt

	for _, n := range nums {
		if max < n {
			max = n
		}

		if min > n {
			min = n
		}
	}

	diff := max - min - (k * 2)

	if diff > 0 {
		return diff
	} else {
		return 0
	}
}
