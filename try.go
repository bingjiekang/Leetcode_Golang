package main

import (
	"fmt"
	// "math"
)

func main() {
	// // fmt.Println(countAndSay(5))
	// a := 3
	// tm := fmt.Sprintf("%b", a)
	fmt.Println(CountBit(1))
	fmt.Println(CountBit(2))
	fmt.Println(CountBit(3))
	fmt.Println(CountBit(7))
	fmt.Println(CountBit(15))
	fmt.Println(CountBit(12))

}

func CountBit(n int) (sum int) {
	for n/2 != 0 || n == 1 {
		if n%2 == 1 {
			sum++
		}
		n /= 2
	}
	return
}
