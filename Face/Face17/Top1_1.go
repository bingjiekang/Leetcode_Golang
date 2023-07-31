package main

// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

// 示例 1：

// 输入: s = "leetcode"
// 输出: false

func isUnique(astr string) bool {
	// // 第一种方法通过map存储，当出现次数大于1时，返回false，如果都不大于1则返回true
	// var hmap map[rune]int = make(map[rune]int,0)
	// for _,v := range astr{
	//     hmap[v]++
	//     if hmap[v]>1{
	//         return false
	//     }
	// }
	// return true

	// 第二种方法利用位运算，通过a-z范围内每个减去a得到可以移动的位数，对1进行左移，例如a得话就是对1进行左移0位，b的话就是对1左移1位，c的话就是对1左移2位，这个数与sult进行与操作，若为0证明他没有出现过，对sult进行或运算，将这个位 置为1，如果出现过就返回false，如果遍历完没有出现过返回true
	// 对sult进行初始复制为0
	var sult rune = int32(0)
	// 遍历astr
	for _, v := range astr {
		// 得到a-z对应移动的位数
		su := int32(v) - 97
		// 对1进行移位操作，再与sult进行与，如果结果为0证明没出现过，进行赋值
		if sult&(1<<su) == 0 {
			sult |= (1 << su)
		} else { // 出现过返回false
			return false
		}
	}
	// 全部没出现过 返回true
	return true

}
