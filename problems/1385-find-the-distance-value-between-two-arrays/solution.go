func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	answer := 0

	sort.Ints(arr2)

	for _, v := range arr1 {
		if binarySearch(arr2, v, d) {
			answer += 1
		}
	}

	return answer
}

func binarySearch(arr []int, n int, distance int) bool {
	l := 0
	h := len(arr) - 1

	for l <= h {
		m := l + (h-l)/2

		if int(math.Abs(float64(arr[m]-n))) <= distance {
			return false
		} else if arr[m] < n {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return true
}
