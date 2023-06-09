package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	tml := []byte{'1'}
	tmr := []byte{'1'}
	var sult []rune = []rune{'1'}
	var length int
	for i := 1; i < n; i++ {
		length = len(sult)
		temp := 1
		for li := 1; li < length; li++ {
			if sult[li] == sult[li-1] {
				temp++
			} else {
				tml = append(tml, byte(temp)+48)
				tmr = append(tmr, byte(sult[li-1]))
				temp = 1
			}
			if li == length-1 {
				tml = append(tml, byte(temp)+48)
				tmr = append(tmr, byte(sult[li]))
			}
		}
		for i, v := range tml {
			sult = append(sult, rune(v))
			sult = append(sult, rune(tmr[i]))
		}
		sult = sult[length:]
		tml = []byte{}
		tmr = []byte{}

	}
	return string(sult)

}
