package main

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

// 示例:

// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.min();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.min();   --> 返回 -2.

type MinStack struct {
	length int
	st     []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var temp MinStack
	temp.st = make([]int, 0)
	return temp
}

func (this *MinStack) Push(x int) {
	this.length++
	this.st = append(this.st, x)
}

func (this *MinStack) Pop() {
	if this.length > 0 {
		this.length--
		this.st = this.st[0:this.length]
	}

}

func (this *MinStack) Top() int {
	return this.st[this.length-1]
}

func (this *MinStack) Min() int {
	var min int = this.st[0]
	for _, v := range this.st {
		if v < min {
			min = v
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
