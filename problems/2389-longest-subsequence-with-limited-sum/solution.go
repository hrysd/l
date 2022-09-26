func answerQueries(nums []int, queries []int) []int {
	answer := make([]int, len(queries))
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	for i, q := range queries {
		answer[i] = binarySearch(nums, q)
	}

	return answer
}

func binarySearch(arr []int, target int) int {
	l := 0
	h := len(arr) - 1
	m := 0

	for l <= h {
		m = (l + h) / 2
		if arr[m] > target {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}

