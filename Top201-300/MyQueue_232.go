package main

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
// 实现 MyQueue 类：
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
// 输入：
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// 输出：
// [null, null, null, 1, 1, false]

type MyQueue struct {
	stack []int
}

func Constructor() (tm MyQueue) {
	return
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	temp := this.stack[0]
	this.stack = this.stack[1:]
	return temp
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	length := len(this.stack)
	if length == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
