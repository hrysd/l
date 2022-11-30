func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	answer := 0

	for _, n := range nums {
		if n == 1 {
			count += 1
		} else {
			if count > answer {
				answer = count
			}

			count = 0
		}
	}

	if count > answer {
		answer = count
	}

	return answer
}
