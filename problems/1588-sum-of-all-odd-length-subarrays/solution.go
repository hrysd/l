func sumOddLengthSubarrays(arr []int) int {
	result := 0
	length := len(arr)

	odds := []int{}

	for i := 1; i <= length; i += 2 {
		odds = append(odds, i)
	}

	for _, lengthOfSubArray := range odds {
		for i := 0; i+lengthOfSubArray <= length; i++ {
			for _, v := range arr[i : i+lengthOfSubArray] {
				result += v
			}
		}
	}

	return result
}
