package main

import (
	"fmt"
)

func main() {
	var n, m int
	// var lt [][]int = []int{}
	var coast, worth []int = []int{}, []int{}
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		var k, t int
		fmt.Scan(&k, &t)
		coast = append(coast, k)
		worth = append(worth, t)
	}
	var sult []int = make([]int, m+1)
	for i := 0; i < n; i++ {
		for j := m; j >= coast[i]; j-- {
			sult[j] = max(sult[j-1], sult[j-coast[i]]+worth[i])
		}
	}
	var mx int
	for i := 0; i < m; i++ {
		if sult[i] > mx {
			mx = sult
		}
	}
	return mx
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
