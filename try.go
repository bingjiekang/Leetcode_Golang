package main

import (
	"fmt"
)

func main() {
	var n, m, cout int
	var mxlt [1e3][1e3]int
	fmt.Scanf("%d %d", &n, &m)
	for j := 1; j <= n; j++ {
		var tm int
		fmt.Scanf("%d", &tm)
		for i := j; i > 0; i-- {
			if i == j {
				mxlt[i][j] = tm
			} else {
				mxlt[i][j] = min(mxlt[i+1][j], mxlt[i][j-1])
			}
		}
	}

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scanf("%d %d", &l, &r)
		mxlt[l][r]--
		if mxlt[l][r] < 0 {
			fmt.Println(cout)
			return
		}
		cout++
	}
	fmt.Println(cout)

}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	m, n = m-1, n-1
	var lt []int = make([]int, 0)
	st := m * n
	var yt, ot int = 3, 0
	for i := 0; i < st; i++ {
		if i < n {
			lt = append(lt, i+2)
		} else if i%n == 0 {
			lt = append(lt, yt)
			yt++
		} else {
			lt = append(lt, lt[i-1]+lt[i%n+n*ot])
			ot++
		}
	}
	return lt[st-1]
}
