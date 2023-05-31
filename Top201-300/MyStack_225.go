package main

// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
// 实现 MyStack 类：
// void push(int x) 将元素 x 压入栈顶。
// int pop() 移除并返回栈顶元素。
// int top() 返回栈顶元素。
// boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

// 方法:通过两个数组实现栈的方法,结构体,定义两个数组,由于栈是后入先出,相当于数组每加进去一个就放在最前面,则实现方法通过现将待加入数据放入另一个数组B,其次再将原数组A的全部数据加入到这个数组B,然后对换两个数组的数据,并将另一个数组B置为空即可
type MyStack struct {
	quen1, quen2 []int
}

func Constructor() (s MyStack) {
	return
}

// 实现入栈功能
func (this *MyStack) Push(x int) {
	this.quen2 = append(this.quen2, x)
	for _, v := range this.quen1 {
		this.quen2 = append(this.quen2, v)
	}
	this.quen2, this.quen1 = this.quen1[0:0], this.quen2
}

// 弹出栈顶元素(也是列表第一个元素)
func (this *MyStack) Pop() int {
	if len(this.quen1) > 0 {
		temp := this.quen1[0]
		this.quen1 = this.quen1[1:]
		return temp
	} else {
		return this.quen2[0]
	}

}

// 查看顶端元素(列表第一个元素)
func (this *MyStack) Top() int {
	return this.quen1[0]
}

// 判断是否为空
func (this *MyStack) Empty() bool {
	if len(this.quen1) > 0 {
		return false
	}
	return true
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
