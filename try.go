package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(5))

	// merge()
}

func generateMatrix(n int) (sult [][]int) {
	sult = make([][]int, 0)
	for i := 0; i < n; i++ {
		var lt []int = make([]int, n)
		sult = append(sult, lt)
	}
	// 1-右，2-下，3-左,4-上
	// var move [4]int = [4]int{1,2,3,4}
	maxSize := n * n
	row, col := 0, 0
	mv := 1
	for value := 1; value <= maxSize; value++ {
		sult[row][col] = value
		switch mv {
		case 1:
			col++
		case 2:
			row++
		case 3:
			col--
		case 4:
			row--
		}
		if col >= n || (mv == 1 && sult[row][col] != 0) {
			mv = 2
			col--
			row++
		} else if row >= n || (mv == 2 && sult[row][col] != 0) {
			mv = 3
			row--
			col--
		} else if col <= -1 || (mv == 3 && sult[row][col] != 0) {
			mv = 4
			col++
			row--
		} else if row <= -1 || (mv == 4 && sult[row][col] != 0) {
			mv = 1
			row++
			col++
		}
	}

	return
}
