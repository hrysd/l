func distinctAverages(nums []int) int {
	m := map[string]int{}

	sort.Ints(nums)

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		a := float64(nums[i]+nums[j]) / 2.0
		s := strconv.FormatFloat(a, 'f', 2, 32)
		m[s] += 1
	}

	return len(m)
}
