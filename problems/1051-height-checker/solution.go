func heightChecker(heights []int) int {
	clone := append([]int{}, heights...)
	sort.Ints(clone)
	count := 0

	for i, v := range heights {
		if v != clone[i] {
			count++
		}
	}

	return count
}
