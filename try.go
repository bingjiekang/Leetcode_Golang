package main

import (
	"fmt"
)

func main() {
	var lt []int = []int{1}
	lt = append(lt, 2)
	lt = append(lt, 3)
	x := append(lt, 4)
	y := append(lt, 5)
	fmt.Println(lt, x, y)
	fmt.Printf("%T %T", &x, &y)
	fmt.Println(&x[3], &y[3])

}
