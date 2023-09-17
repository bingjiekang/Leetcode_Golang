package main

import (
	"fmt"
)

func main() {
	fmt.Println(backspaceCompare("ab##", "c#d#"))

}

func backspaceCompare(s string, t string) bool {
	var lts, ltt []rune = make([]rune, 0), make([]rune, 0)
	var couts, coutt int
	for _, v := range s {
		if v == '#' && couts > 0 {
			lts = lts[0 : couts-1]
			couts--
		} else if v != '#' {
			lts = append(lts, v)
			couts++
		}
	}

	for _, v := range t {
		if v == '#' && coutt > 0 {
			ltt = ltt[0 : coutt-1]
			coutt--
		} else if v != '#' {
			ltt = append(ltt, v)
			coutt++
		}
	}
	if string(lts) == string(ltt) {
		return true
	} else {
		return false
	}

}
