func largestGoodInteger(num string) string {
	max := math.MinInt

	for i := 2; i < len(num); i++ {
		if num[i-2] != num[i-1] || num[i-1] != num[i] {
			continue
		}

		n, _ := strconv.Atoi(num[i-2 : i+1])

		if n > max {
			max = n
		}
	}

	if max > math.MinInt {
		return fmt.Sprintf("%03d", max)
	} else {
		return ""
	}
}
