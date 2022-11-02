func slowestKey(releaseTimes []int, keysPressed string) byte {
	max := math.MinInt
	key := 'a'

	for i := 0; i < len(releaseTimes); i++ {
		t := releaseTimes[i]

		if i > 0 {
			t -= releaseTimes[i-1]
		}

		if t > max {
			max = t
			key = rune(keysPressed[i])
		} else if t == max {
			if int(keysPressed[i]-'a') > int(key-'a') {
				key = rune(keysPressed[i])
			}
		}
	}

	return byte(key)
}
