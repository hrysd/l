// https://leetcode.com/problems/number-complement/discuss/96103/maybe-fewest-operations
func findComplement(num int) int {
	mask := num

	mask = mask>>1 | mask
	mask = mask>>2 | mask
	mask = mask>>4 | mask
	return num ^ mask
}
