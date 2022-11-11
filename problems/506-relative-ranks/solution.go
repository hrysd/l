func findRelativeRanks(score []int) []string {
	ranks := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}

	m := map[int]string{}

	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	for i, v := range sorted {
		if rank, ok := ranks[i+1]; ok {
			m[v] = rank
		} else {
			m[v] = strconv.Itoa(i + 1)
		}
	}

	answer := make([]string, len(score))

	for i, v := range score {
		answer[i] = m[v]
	}

	return answer
}
