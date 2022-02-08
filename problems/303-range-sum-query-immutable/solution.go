type NumArray struct {
	Body []int
}

func Constructor(nums []int) NumArray {
	return NumArray{Body: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int

	for i, v := range this.Body {
		if left <= i && i <= right {
			sum += v
		}
	}

	return sum
}
