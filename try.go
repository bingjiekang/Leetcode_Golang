package main

import (
	"fmt"
)

func main() {
	// price := 2.3
	// value := 33000
	// tempAmount := float64(price / 33000 * float64(value))
	tempAmount := 3.44132
	tm := 0.004
	tmp := fmt.Sprintf("%.2f", tempAmount+tm)
	// totalAmount := tempAmount * 1e2 * 1e-2
	fmt.Println(tempAmount, tmp)
}
