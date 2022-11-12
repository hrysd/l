func averageValue(nums []int) int {
	sum := 0
	count := 0

	for _, n := range nums {
		if n%6 == 0 {
			sum += n
			count += 1
		}
	}

	if sum == 0 {
		return 0
	} else {
		return sum / count
	}
}
