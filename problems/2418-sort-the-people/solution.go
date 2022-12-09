func sortPeople(names []string, heights []int) []string {
	m := map[int]string{}

	for i, h := range heights {
		m[h] = names[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(heights)))

	answer := []string{}
	for _, h := range heights {
		answer = append(answer, m[h])
	}
	return answer
}
