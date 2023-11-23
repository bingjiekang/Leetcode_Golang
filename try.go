package main

import (
	"fmt"
	"sort"
)

func main() {
	// var tmp []int = []int{4,3,12,5,6}
	// sort.Reverse()
	// sort.Interface
}

func merge(intervals [][]int) (sult [][]int) {
	var newNums []Nums
	var temp Nums
	for _, v := range intervals {
		temp.Start = v[0]
		temp.End = v[1]
		newNums = append(newNums, temp)
	}
	sort.Sort(md(newNums))
	fmt.Println(newNums)
}

type Nums struct {
	Start int
	End   int
}

type md []Nums

func (m md) Len() int {
	return len(m)
}

func (m md) Less(i, j int) bool {
	return m[i].Start < m[j].Start
}

func (m md) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
