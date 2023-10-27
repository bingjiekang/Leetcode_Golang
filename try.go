package main

import (
	"fmt"
	"sort"
)

func main() {
	var tm []int = []int{2, 4, 1, 5, 6, 9, 5}
	sort.Ints(tm)
	// // var tm []string = []string{"acd", "bed", "accd"}
	// // fmt.Println(minNumBooths(tm))
	// var dmm, dmd map[rune]int = make(map[rune]int, 0), make(map[rune]int, 0)
	// dmm['r'] = 1
	// dmd['t'] = 3
	// dmm['r'] = dmd['t']
	// fmt.Println(dmm['r'], dmm['t'], dmm['u'])
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	lentap, lendnk := len(staple), len(drinks)
	var sult int
	sort.Ints(staple)
	sort.Ints(drinks)
	for i := lentap - 1; i >= 0; i-- {
		for j := lendnk - 1; j >= 0; j-- {
			if staple[i]+drinks[j] <= x {
				sult = ((i + 1) * (j + 1)) % (1e9 + 7)
				return sult
			}
		}
	}
	return sult
}
