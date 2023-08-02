package main

import (
	"fmt"
	"strconv"
	"strings"
)

// "strconv"
// "time"
// "math"

func main() {
	// var hash map[string]int = make(map[string]int, 0)
	// n, m := 1, 2
	// temp := strconv.Itoa(n) + "," + strconv.Itoa(m)
	// fmt.Println(temp)
	// hash[temp] = 1
	// fmt.Println(hash)
	// // var su []string = make([]string, 0)
	// su := strings.Split(temp, ",")
	// t1, _ := strconv.ParseInt(su[0], 10, 64)
	// t2, _ := strconv.ParseInt(su[0], 10, 64)
	// var tm [4][7]int = [4][7]int{

	// 	[14, 15, 100, 9, 3],
	// 	[32, 1, 9, 3, 5],
	// 	[15, 29, 2, 6, 8, 7],
	// 	[7, 10]

	// }

}

func computeSimilarities(docs [][]int) (sult []string) {
	// 用map型数组存储每组数组对应的组数数组
	length := len(docs)
	var curd map[int][]int = make(map[int][]int, 0)
	for i, v := range docs {
		for _, w := range v {
			// 判断数w是否已经在map1里
			_, ok := curd[w]
			if !ok { // 如果不在则声明一个
				curd[w] = make([]int, length+1, length+2)
			}
			// 表示这个数所在的组为1
			curd[w][i] = 1
			// 最后一个数表示这个数对应组里有几个数
			curd[w][length]++
		}
	}
	// 定义一个map2用来存储对应两组出现的次数,表示交集
	// var hash map[string]int = make(map[string]int, 0)
	// 定义一个二维切片存储对应的次数
	var lt [][]int = make([][]int, length, length+1)
	// 遍历每个map1,提取对应两组出现的组数 对其进行加1表示交集数
	for _, v := range curd {
		if v[length] > 1 {
			for i := 0; i < length-1; i++ {
				for j := i + 1; j < length; j++ {
					lt[i] = make([]int, length, length+1)
					if v[i] == 1 && v[j] == 1 {
						lt[i][j]++
					}
				}
			}
		}
	}
	clear(curd)
	// 遍历每个数组获取对应数组下标的长度 相加-交集即为并集,想除并转换加入即可
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			// 求对应并集
			ad := len(docs[i]) + len(docs[j]) - lt[i][j]
			ku := float64(lt[i][j])/float64(ad) + 1e-9
			tsu := fmt.Sprintf("%d,%d: %.4f", i, j, ku)
			sult = append(sult, tsu)
		}
	}
	return
}
