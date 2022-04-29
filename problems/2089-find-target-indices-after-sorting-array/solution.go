func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	result := []int{}

	for i, v := range nums {
		if v == target {
			result = append(result, i)
		}

	}

	return result
}
