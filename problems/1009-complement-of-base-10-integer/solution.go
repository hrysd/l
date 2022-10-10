func bitwiseComplement(n int) int {
	return (1<<len(strconv.FormatInt(int64(n), 2)) - 1) ^ n
}
