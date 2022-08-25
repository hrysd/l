func sortEvenOdd(nums []int) []int {
	odd := []int{}
	even := []int{}

	for i, n := range nums {
		if i%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(odd)))
	sort.Ints(even)

	result := []int{}
	for i := 0; i < len(nums); i++ {
		var n int

		if i%2 == 0 {
			n = even[0]
			even = even[1:]
		} else {
			n = odd[0]
			odd = odd[1:]
		}

		result = append(result, n)
	}

	return result
}
