func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1 := count(nums1)
	m2 := count(nums2)
	m3 := count(nums3)

	result := []int{}

	for i := 1; i <= 100; i++ {
		n1, _ := m1[i]
		n2, _ := m2[i]
		n3, _ := m3[i]

		if n1+n2+n3 > 1 {
			result = append(result, i)
		}
	}

	return result
}

func count(nums []int) map[int]int {
	m := map[int]int{}

	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 1
		}
	}

	return m
}
