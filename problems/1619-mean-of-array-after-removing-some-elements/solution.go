func trimMean(arr []int) float64 {
	sort.Ints(arr)

	l := len(arr)

	sum := 0
	count := 0

	for i := l / 20; i < l-(l/20); i++ {
		sum += arr[i]
		count++
	}

	return float64(sum) / float64(count)
}
