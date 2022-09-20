func findKDistantIndices(nums []int, key int, k int) []int {
	answer := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			d := int(math.Abs(float64(i - j)))

			if d <= k && nums[j] == key {
				answer = append(answer, i)
				break
			}
		}
	}

	return answer
}
