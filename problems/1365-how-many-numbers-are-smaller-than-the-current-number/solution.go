func smallerNumbersThanCurrent(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	m := map[int]int{}
	result := make([]int, len(nums))

	advanced := 0
	for _, v := range sorted {
		_, ok := m[v]

		if ok {
			advanced++
		} else {
			m[v] = len(m) + advanced

		}
	}

	for i, v := range nums {
		result[i] = m[v]
	}

	return result
}
