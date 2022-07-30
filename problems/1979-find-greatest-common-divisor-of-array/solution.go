func findGCD(nums []int) int {
	max := nums[0]
	min := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return gcd(min, max)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

