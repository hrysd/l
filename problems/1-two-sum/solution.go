func twoSum(nums []int, target int) []int {
	result := make([]int, 2)

Exit:
	for i, num := range nums {
		for anotherIndex, anotherNum := range nums {
			if i == anotherIndex {
				continue
			}

			isSame := num+anotherNum == target

			if isSame {
				result[0] = i
				result[1] = anotherIndex
				break Exit
			}
		}
	}

	return result
}
