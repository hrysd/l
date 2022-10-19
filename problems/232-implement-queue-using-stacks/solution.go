type MyQueue struct {
	Queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		Queue: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyQueue) Pop() int {
	n := this.Queue[0]
	this.Queue = this.Queue[1:]
	return n
}

func (this *MyQueue) Peek() int {
	return this.Queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.Queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
