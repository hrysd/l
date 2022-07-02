func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	results := [][]int{}
	min := math.MaxInt

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		if min == diff {
			results = append(results, []int{arr[i-1], arr[i]})
		} else if min > diff {
			min = diff
			results = [][]int{{arr[i-1], arr[i]}}
		}

	}

	return results
}
