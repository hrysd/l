func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := map[int]int{}

	for _, v := range arr1 {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	result := []int{}

	for _, v := range arr2 {
		count, _ := m[v]

		for i := 0; i < count; i++ {
			result = append(result, v)
		}

		delete(m, v)
	}

	remains := []int{}

	for k, count := range m {
		for i := 0; i < count; i++ {
			remains = append(remains, k)
		}
	}

	sort.Ints(remains)

	return append(result, remains...)
}
