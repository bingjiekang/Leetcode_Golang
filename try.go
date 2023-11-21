package main

import (
	"fmt"
	"strings"
)

func main() {
	// price := 2.3
	// value := 33000
	// tempAmount := float64(price / 33000 * float64(value))
	// tempAmount := 3.44132
	// tm := 0.004
	// tmp := fmt.Sprintf("%.2f", tempAmount+tm)
	// // totalAmount := tempAmount * 1e2 * 1e-2
	fmt.Println(findWords([]string{"omk"}))
}

var strand []string = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

func findWords(words []string) (sult []string) {
	var dict map[rune]int = make(map[rune]int, 0)
	for k, v := range strand {
		for _, t := range v {
			dict[t] = k
		}
	}
	for _, v := range words {
		// value,length := 0,len(v)
		vt := strings.ToLower(v)
		target := dict[rune(vt[0])]
		success := true
		for _, tmp := range vt {
			if target != dict[tmp] {
				success = false
				break
			}
		}
		if success {
			sult = append(sult, vt)
		}
	}
	return sult
}
