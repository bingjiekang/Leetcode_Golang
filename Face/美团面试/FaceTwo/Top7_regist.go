// 7. 美团商家注册系统 mid
// 请你开发一个美团商家测试系统，并用等价划分法确认商家注册信息是否成功。
// 商家信息必须满足以下条件：
// 1. 系统中第一次注册的商家名字，被视为主店。
// 2. 系统中若出现重名商家，需要判断地址是否已存在该商家。若存在，则注册失败。否则注册成功，该商家被视为分店。
// 3. 商家的名字和地址必须由小写的英文字母组成，否则注册失败。
// 请你输出每个商家的信息，按商家名字的字典序升序输出。需要输出商家名字，商家主店地址，商家分店数量。

// 输入描述：
// 第一行输入一个正整数n，代表注册信息数量。
// 接下来的n行，每行输入两个字符串，用空格隔开。分别代表商家名字和商家地址。
// 1<= n <= 1000
// 给定的商家名字和商家地址字符串长度不超过 20，且不包含空格。
// 输出描述：
// 按商家名字字典序输出全部商家信息。每行输出一个，分别输出商家名字，商家主店地址，商家分店数量，用空格隔开。
// 示例1
// 输入例子：
// 5
// ranko mt
// ranko op
// ranko op
// Ranko ok
// red ok
// 输出例子：
// ranko mt 1
// red ok 0

// 通过切片字典记录出现的店铺名称和对应的地址,如果出现相同的地址,则不计入,如果不同则为分店,如果没有店铺名则是主店,切片字典第一个即为主店
package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	// 存储对应店面和对应主店和分店地址,第一个为主店,其他为分店
	dt := map[string][]string{}
	// 用来存储并遍历店铺名称,按字典序排列
	tgname := []string{}
	fmt.Scanf("%d", &m)

	for i := 0; i < m; i++ {
		var name, add string

		fmt.Scanln(&name, &add)
		// 如果输入的店铺和地址不为小写则不计入,结束后续操作
		if !pdown(name) || !pdown(add) {
			continue
		}
		// 判断是否已经存在
		_, ok := dt[name]
		if !ok { // 若不存在,则加入店铺和对应地址,以及店铺名在切片中的信息
			dt[name] = append(dt[name], add)
			tgname = append(tgname, name)

		} else { // 若已经存在 则判断新的店铺地址是否和已经存在的店铺地址冲突
			var tag bool = true
			for _, v := range dt[name] {
				if add == v { // 如果冲突则不加入
					tag = false
					break
				}
			}
			if tag { // 如果不冲突则加入这个新地址
				dt[name] = append(dt[name], add)

			}
		}
	}
	// 对店铺名进行字典序排列
	sort.Strings(tgname)
	// 输出对应信息即可
	for _, v := range tgname {
		fmt.Println(v, dt[v][0], len(dt[v])-1)
	}

}

// 用来判断是否都是小写字母
func pdown(nums string) bool {
	for _, v := range nums {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}
