func createTargetArray(nums []int, index []int) []int {
	array := []int{}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		pos := index[i]

		array = append(array[:pos], append([]int{n}, array[pos:]...)...)
	}

	return array
}
