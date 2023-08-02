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
	var hash map[string]int = make(map[string]int, 0)
	// 遍历每个map1,提取对应两组出现的组数 对其进行加1表示交集数
	for _, v := range curd {
		if v[length] > 1 {
			for i := 0; i < length-1; i++ {
				for j := i + 1; j < length; j++ {
					if v[i] == 1 && v[j] == 1 {
						temp := strconv.Itoa(i) + "," + strconv.Itoa(j)
						hash[temp]++
					}
				}
			}
		}
	}
	// 遍历每个map2获取对应数组的长度 相加-交集即为并集,想除并转换加入即可
	for k, v := range hash {
		var tsu string
		su := strings.Split(k, ",")
		t1, _ := strconv.ParseInt(su[0], 10, 64)
		t2, _ := strconv.ParseInt(su[1], 10, 64)
		ad := len(docs[t1]) + len(docs[t2]) - v
		ku := float64(v)/float64(ad) + 1e-9
		if ku > 0 {
			tsu = k + ": " + fmt.Sprintf("%.4f", ku)
		}
		sult = append(sult, tsu)
	}
	return
}
