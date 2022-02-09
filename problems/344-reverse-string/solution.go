import "reflect"

func reverseString(s []byte) {
	swappable := reflect.Swapper(s)
	max := len(s) - 1

	for i, j := 0, max; i < j; i, j = i+1, j-1 {
		swappable(i, j)
	}
}
