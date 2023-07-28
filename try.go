package main

import (
	"fmt"
	"strconv"
)

// "strconv"
// "time"
// "math"

func main() {

	// n := 3

	// for i := 0; i <= 4; i++ {
	// 	switch i {
	// 	case 3:
	// 		fmt.Println(3)
	// 	default:
	// 		fmt.Println(99999)
	// 	}

	// }
	// fmt.Println("end")
	// var lt stack
	// fmt.Println(lt)
	// lt.push(4)
	// fmt.Println(lt)
	// lt.push(6)
	// fmt.Println(lt)
	// fmt.Println(lt.pop())
	// fmt.Println(lt)
	// tm := []rune("123")
	// fmt.Println(tm, int())
	// fmt.Println(strconv.ParseInt("123",10,32).type())
}

func evalRPN(tokens []string) int {
	var lt stack
	for _, v := range tokens {
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
		default:
			temp, _ := strconv.ParseInt(v, 10, 64)
			lt.push(int(temp))
		}
	}
	return lt.pop()

}

// 实现一个堆
type stack struct {
	List []int
}

func (this *stack) push(dig int) {
	this.List = append(this.List, dig)
}

func (this *stack) pop() int {
	length := len(this.List)
	wait := this.List[length-1]
	this.List = this.List[:length-1]
	return wait
}

func (this *stack) isnil() bool {
	return len(this.List) == 0
}
