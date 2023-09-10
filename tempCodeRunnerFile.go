package main

import (
	"fmt"
)

func main() {
	var n, m, cout int
	var sult []int = make([]int, 0)
	fmt.Scanln(&n, &m)
	var tmp int = 1
	var op int
	for op = 1; 1+(op*(op-1)/2) <= m; op++ {
	}
	if op >= n {
		for i := 0; i < n; i++ {
			cout++
			sult = append(sult, tmp)
			tmp += cout
		}
	} else {
		op = n - op
		for i := 0; i < n; i++ {
			cout++
			sult = append(sult, tmp)
			begin := tmp
			tmp += cout
			for j := begin + 1; j < tmp && j < m && op > 0; j++ {
				sult = append(sult, tmp)
				op--
				i++
			}

		}
	}
	for i, v := range sult {
		if i != n-1 {
			fmt.Printf("%d ", v)
		} else {
			fmt.Printf("%d", v)
		}

	}
}
