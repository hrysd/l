func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	max := len(nums) - 1

	return (nums[max] * nums[max-1]) - (nums[1] * nums[0])
}
