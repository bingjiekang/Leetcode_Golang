package main

import (
	"strconv"
)

// 给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。

// 请你计算该表达式。返回一个表示表达式值的整数。

// 注意：
// 有效的算符为 '+'、'-'、'*' 和 '/' 。
// 每个操作数（运算对象）都可以是一个整数或者另一个表达式。
// 两个整数之间的除法总是 向零截断 。
// 表达式中不含除零运算。
// 输入是一个根据逆波兰表示法表示的算术表达式。
// 答案及所有中间计算结果可以用 32 位 整数表示。

// 解析:通过数组实现一个堆,然后对数据进行分析,如果是+-*/其中之一,则进行数据提取和运算,如果为字符串数字,则进行字符串转换成数字,进行运算既可

func evalRPN(tokens []string) int {
	var lt stack
	for _, v := range tokens {
		// 如果为运算符则通过提取前两个数,进行运算
		switch v {
		case "+":
			num2 := lt.pop()
			num1 := lt.pop()
			sum := num1 + num2
			lt.push(sum)
		case "-":
			num2 := lt.pop()
			num1 := lt.pop()
			sum := num1 - num2
			lt.push(sum)
		case "*":
			num2 := lt.pop()
			num1 := lt.pop()
			sum := num1 * num2
			lt.push(sum)
		case "/":
			num2 := lt.pop()
			num1 := lt.pop()
			sum := num1 / num2
			lt.push(sum)
		default: // 否则则对数字进行转换后加入到堆中
			temp, _ := strconv.ParseInt(v, 10, 64)
			lt.push(int(temp))
		}
	}
	return lt.pop() // //最后仅剩的即为结果

}

// 实现一个堆
type stack struct {
	List []int
}

// 压入堆中一个数据
func (this *stack) push(dig int) {
	this.List = append(this.List, dig)
}

// 弹出顶层数据
func (this *stack) pop() int {
	length := len(this.List)
	wait := this.List[length-1]
	this.List = this.List[:length-1]
	return wait
}

// 判断堆是否为空
func (this *stack) isnil() bool {
	return len(this.List) == 0
}
