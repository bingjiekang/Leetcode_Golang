package main

import (
	"fmt"
)

func main() {
	m := Constructor()
	m.Push(3)
	m.Push(4)
	m.Push(5)
	fmt.Println(m.quen1, m.quen2)
	fmt.Println(m.Pop())
	fmt.Println(m.quen1, m.quen2)
}

type MyStack struct {
	quen1, quen2 []int
}

func Constructor() (s MyStack) {
	return
}

func (this *MyStack) Push(x int) {
	this.quen2 = append(this.quen2, x)
	for _, v := range this.quen1 {
		this.quen2 = append(this.quen2, v)
	}
	this.quen2, this.quen1 = this.quen1[0:0], this.quen2
}

func (this *MyStack) Pop() int {
	temp := this.quen1[0]
	this.quen1 = this.quen1[1:]
	return temp
}
