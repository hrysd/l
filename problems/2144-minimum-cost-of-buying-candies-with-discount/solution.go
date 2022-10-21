func minimumCost(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))

	answer := 0
	for i, c := range cost {
		if (i+1)%3 != 0 {
			answer += c
		}
	}

	return answer
}
