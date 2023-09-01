package main

import (
	"strconv"
	"strings"
)

func main() {
	// var a []int = make([]int, 3)
	// v := []int{1, 2, 3}
	// fmt.Println(a, v)
	// copy(a, v)
	// var tm map[[]int]bool = make(map[interface{}]bool)
	// fmt.Println(a, v)

	// tm[[1,2,1]] = true
	// fmt.Println(permute([]int{1, 2, 3}))
}

func compareVersion(version1 string, version2 string) int {
	lt1, lt2 := strings.Split(version1, "."), strings.Split(version2, ".")
	length1, length2 := len(lt1), len(lt2)
	for i := 0; i < length1 || i < length2; i++ {
		x, y := 0, 0
		if i < length1 {
			x, _ = strconv.Atoi(lt1[i])
		}
		if i < length2 {
			y, _ = strconv.Atoi(lt2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}

	}
	return 0
}
