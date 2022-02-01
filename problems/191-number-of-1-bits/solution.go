func hammingWeight(num uint32) int {
	var count int
	n := num

	for n != 0 {
		bit := n & 1

		if bit == 1 {
			count = count + 1
		}

		n = n >> 1
	}

	return count
}
