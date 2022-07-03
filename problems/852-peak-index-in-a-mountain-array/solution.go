func peakIndexInMountainArray(arr []int) int {
	l := len(arr)
	left, right := 0, l-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
