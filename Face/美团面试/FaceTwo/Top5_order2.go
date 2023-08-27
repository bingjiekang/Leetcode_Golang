// 5. 小美的数组操作2 easy
// 小美拿到了一个数组，她每次可以进行如下操作：
// 选择两个元素，一个加 1，另一个减 1。
// 小美总共进行了k次操作。她希望你回答最终数组是否是非降序，你能帮帮她吗？
// 请注意，元素可能会被减成负数！

// 输入描述：
// 第一行输入一个正整数 t，代表询问次数。
// 每次询问首先第一行输入两个正整数n和k，代表数组长度和操作次数。
// 接下来的一行输入n个正整数 a_i，代表初始数组。
// 接下来的k行，每行输入两个正整数u,v，代表使得第u个元素加 1，第v个元素减 1。
// 1<= t,n,k,ai <=  100
// 输出描述：
// 输出t行，每行输出该次询问的答案。
// 如果数组变成了非降序，则输出"Yes"。否则输出 "No"。

// 示例1
// 输入例子：
// 2
// 3 2
// 3 4 5
// 2 3
// 1 2
// 3 2
// 3 4 5
// 2 3
// 2 3
// 输出例子：
// Yes
// No
// 例子说明：
// 第一组询问，操作两次后数组变成[4,4,4]，为非降序。
// 第二组询问，操作两次后数组变成[3,6,3]，并不是非降序。

// 通过对指定位进行操作,然后得到切片结果,最后判断是否是非降序即可

package main

import (
	"fmt"
)

func main() {
	// 处理输入
	var m int
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		var ln, cout int
		lst := []int{}
		// 获得输入的数组长度和操作次数
		fmt.Scanln(&ln, &cout)
		var temp int
		// 将数据加入到对应数组中
		for j := 0; j < ln; j++ {
			fmt.Scanf("%d", &temp)
			lst = append(lst, temp)
		}
		// 遍历操作次数
		for t := 0; t < cout; t++ {
			var tlf, trg int
			fmt.Scanln(&tlf, &trg)
			if tlf <= ln && trg <= ln {
				lst[tlf-1]++
				lst[trg-1]--
			}
		}
		// 如果是非降序则返回yes否则返回no
		if PdBool(lst) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	}
}

// 实现一个函数，确定这个是否是非降序
func PdBool(num []int) bool {
	length := len(num)
	for i := 1; i < length; i++ {
		if num[i] < num[i-1] {
			return false
		}
	}
	return true
}
