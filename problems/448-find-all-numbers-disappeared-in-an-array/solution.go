func findDisappearedNumbers(nums []int) []int {
	memo := map[int]int{}

	for _, n := range nums {
		memo[n] = 1
	}

	answer := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := memo[i]; !ok {
			answer = append(answer, i)
		}
	}

	return answer
}
