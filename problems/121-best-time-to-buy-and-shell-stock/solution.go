func maxProfit(prices []int) int {
	min, max := prices[0], 0

	for i, price := range prices {
		if min > price {
			min = price
		}

		profit := prices[i] - min

		if profit > max {
			max = profit
		}
	}

	return max
}
