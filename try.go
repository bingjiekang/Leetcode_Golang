package main

import "fmt"

func main() {

	var lt [5]int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&lt[i])
	}
	fmt.Println(lt)
}
