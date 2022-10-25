func mostFrequent(nums []int, key int) int {
	m := map[int]int{}
	var answer int

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == key {
			m[nums[i]] += 1

			if m[nums[i]] > m[answer] {
				answer = nums[i]
			}
		}
	}

	return answer
}
