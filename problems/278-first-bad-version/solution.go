func firstBadVersion(n int) int {
	start := 1
	end := n

	for start < end {
		version := (start + end) / 2

		isBad := isBadVersion(version)

		if isBad {
			end = version
		} else {
			start = version + 1
		}
	}

	return end
}
