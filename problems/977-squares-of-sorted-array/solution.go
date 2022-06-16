func sortedSquares(nums []int) []int {
	result := []int{}

	i, j := 0, len(nums)-1

	for i <= j {
		left := int(math.Abs(float64(nums[i])))
		right := int(math.Abs(float64(nums[j])))

		if left > right {
			result = append([]int{left * left}, result...)
			i++
		} else {
			result = append([]int{right * right}, result...)
			j--
		}
	}

	return result
}
