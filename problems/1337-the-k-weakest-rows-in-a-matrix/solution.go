type Pair struct {
	p, i int
}

func kWeakestRows(mat [][]int, k int) []int {
	memo := make([]Pair, len(mat))
	for i, m := range mat {
		p := search(m)

		memo[i] = Pair{p, i}
	}

	sort.SliceStable(memo, func(i, j int) bool { return memo[i].p < memo[j].p })

	result := make([]int, k)

	for i := 0; i < k; i++ {
		result[i] = memo[i].i
	}
	return result
}

func search(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
