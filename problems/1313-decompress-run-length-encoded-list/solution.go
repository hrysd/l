func decompressRLElist(nums []int) []int {
	loop := len(nums) / 2
	result := []int{}

	for i := 0; i < loop; i++ {
		index := 2 * i
		freq := nums[index]
		val := nums[index+1]

		for j := 0; j < freq; j++ {
			result = append(result, val)
		}
	}

	return result
}
