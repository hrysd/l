func countPairs(nums []int, k int) int {
	count := 0
	indicesMap := map[int][]int{}

	for i, v := range nums {
		value, exist := indicesMap[v]

		if exist {
			indicesMap[v] = append(value, i)
		} else {
			indicesMap[v] = []int{i}
		}
	}

	for _, indices := range indicesMap {
		for i := 0; i < len(indices); i++ {
			for j := 1; j < len(indices); j++ {
				if i < j {
					if (indices[i]*indices[j])%k == 0 {
						count++
					}

				}
			}
		}
	}

	return count
}
