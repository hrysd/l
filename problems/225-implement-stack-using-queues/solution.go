type MyStack struct {
	Stack []int
}

func Constructor() MyStack {
	return MyStack{Stack: []int{}}
}

func (this *MyStack) Push(x int) {
	this.Stack = append([]int{x}, this.Stack...)
}

func (this *MyStack) Pop() int {
	pop := this.Stack[0]
	this.Stack = this.Stack[1:]
	return pop
}

func (this *MyStack) Top() int {
	return this.Stack[0]
}

func (this *MyStack) Empty() bool {
	return len(this.Stack) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
