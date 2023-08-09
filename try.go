package main

type MinStack struct {
	length int
	min    int
	st     []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var temp MinStack
	return temp
}

func (this *MinStack) Push(x int) {
	this.length++
	if x < this.min {
		this.min = x
	}
	this.st = make([]int, 0)
	this.st = append(this.st, x)
}

func (this *MinStack) Pop() {
	this.length--
	this.st = this.st[0:this.length]
}

func (this *MinStack) Top() int {
	return this.st[this.length-1]
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
