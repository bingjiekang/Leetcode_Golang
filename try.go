package main

import (
	"fmt"
)

// "strconv"
// "time"
// "math"

func main() {

	fmt.Println((1 + 1) / 2)

}

func canConstruct(ransomNote string, magazine string) bool {
	var hash map[rune]int = make(map[rune]int, 0)
	for _, v := range magazine {
		hash[v]++
	}
	for _, v := range ransomNote {
		// _,err := hash[v]
		if hash[v] > 0 {
			hash[v]--
		} else {
			return false
		}
	}
	return true
}
