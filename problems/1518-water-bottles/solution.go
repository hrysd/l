func numWaterBottles(numBottles int, numExchange int) int {
	answer := numBottles
	for numBottles >= numExchange {
		full := numBottles / numExchange
		empty := numBottles % numExchange
		answer += full

		numBottles = full + empty
	}

	return answer
}
