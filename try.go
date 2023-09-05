package main

import (
	"fmt"
)

func main() {
	fmt.Println(^2, ^4, ^8, ^10, ^-2, ^-4, ^-8, ^-10)
	// var m, temp int
	// var lt []int = make([]int, 0)
	// fmt.Scanln(&m)
	// for i := 0; i < m; i++ {
	// 	fmt.Scanf("%d", &temp)
	// 	lt = append(lt, temp)
	// }
	// sort.Ints(lt)
	// var cout int
	// var dif map[int]int = make(map[int]int, 0)
	// for i := 0; i < m; i++ {
	// 	for lt[i] == -1 && i < m {
	// 		i++
	// 	}

	// 	for j := i + 1; j < m; {
	// 		if (lt[i] == lt[j] || lt[j] == -1) && j < m {
	// 			j++
	// 		} else {
	// 			_, ok := dif[j]
	// 			if !ok {
	// 				lt[j] = -1
	// 				lt[i] = -1
	// 				cout++
	// 				dif[lt[i]] = lt[j]
	// 				break
	// 			} else {
	// 				if dif[lt[i]] != lt[j] {
	// 					dif[lt[i]] = lt[j]
	// 					cout++
	// 					lt[j] = -1
	// 					lt[i] = -1
	// 				}
	// 			}
	// 		}
	// 	}

	// }
	// fmt.Println(cout)

}

func masterMind(solution string, guess string) []int {
	var solu, gues map[byte]int = make(map[byte]int, 0), make(map[byte]int, 0)
	var coutgues_t, coutgues_f int
	for i := 0; i < 4; i++ {
		solu[solution[i]]++
		gues[guess[i]]++
		if solution[i] == guess[i] {
			coutgues_t++
			solu[solution[i]]--
			gues[guess[i]]--
		} else {
			if solu[solution[i]] > 0 && gues[guess[i]] > 0 {
				coutgues_f++
				solu[solution[i]]--
				gues[guess[i]]--
			}
		}
	}
	return []int{coutgues_t, coutgues_f}
}
