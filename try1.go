package main

import (
	"fmt"
)

func main() {
	var n, t, h int
	var t1, h1, happy int
	var lt [][]int = make([][]int, 0)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d%d", &t, &h)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d%d", &t1, &h1, &happy)
		temp := []int{t1, h1, happy}
		lt = append(lt, temp)
	}
	// fmt.Println(lt)
	var sult int
	var stmp []int = []int{0, 0, 0}
	for i, v := range lt {
		stmp = max(stmp+v, v)
	}

}

func max(n, m []int) []int {

}
