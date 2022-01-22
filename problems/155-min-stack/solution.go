type MinStack struct {
	Stack []int
}

func Constructor() MinStack {
	return MinStack{Stack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.Stack = append([]int{val}, this.Stack...)
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[1:]
}

func (this *MinStack) Top() int {
	return this.Stack[0]
}

func (this *MinStack) GetMin() int {
	var min int

	for i, v := range this.Stack {
		if i == 0 {
			min = v
			continue
		}

		if min > v {
			min = v
		}
	}
	return min
}
