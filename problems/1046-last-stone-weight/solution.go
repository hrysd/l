func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones) // ><

		l := len(stones)

		a := stones[l-1]
		stones = stones[:l-1]
		b := stones[l-2]
		stones = stones[:l-2]

		var s int
		if a > b {
			s = a - b
		} else {
			s = b - a
		}

		stones = append(stones, s)
	}

	return stones[0]
}
