package main

import (
	"fmt"
)

func main() {
	newData := []string{"1", "2", "3", "4", "5", "6", "7"}
	for i := 0; i < len(newData); i++ {
		newData = append(newData[:i], newData[i+1:]...)
		fmt.Println(newData)
		i--
	}

}

// func cutSquares(square1 []int, square2 []int) []float64 {
// 	// 先求两个正方形的中心，然后求直线，直线与正方形的交点就是对应的结果
// 	sqrx1, sqry1 := square1[0]+square1[2]/2, square1[1]+square1[2]/2

// }
