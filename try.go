package main

import (
	"fmt"
)

func main() {
	// fmt.Println(backspaceCompare("ab##", "c#d#"))

}

func numTrees(n int) int {
	var G []int = make([]int, n+1)
	G[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
