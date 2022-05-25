func hammingDistance(x int, y int) int {
	xor := x ^ y

	return bits.OnesCount(uint(xor))
}
