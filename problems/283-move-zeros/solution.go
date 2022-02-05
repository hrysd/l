func moveZeroes(nums []int) {
	length := len(nums)
	i, end := 0, length-1

	for i < length && end > 0 {
		n := nums[i]

		if n == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, n)
		} else {
			i++
		}
		end--
	}
}
