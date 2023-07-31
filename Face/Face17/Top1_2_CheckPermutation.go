package main

// 给定两个由小写字母组成的字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

// 示例 1：

// 输入: s1 = "abc", s2 = "bca"
// 输出: true

// 解析:先对字符串长度对比,若不同则返回false,否则返回true,通过map进行遍历和存储即可,先遍历一个字符串,对应字符加1,然后再遍历另外一个字符串,如果出现则减1,当小于1或者不存在时返回false,最后全部都对上则返回true

func CheckPermutation(s1 string, s2 string) bool {
	// 获取字符串长度
	length1, length2 := len(s1), len(s2)
	// 不同则返回false
	if length1 != length2 {
		return false
	} else {
		// 定义一个map用来存储字符出现的次数
		var temp1 map[rune]int = make(map[rune]int, 0)
		// 对出现的字符次数加1
		for _, v := range s1 {
			temp1[v]++
		}
		// 遍历字符串2
		for _, v := range s2 {
			// 得到字符串2中的字符是否存在temp1中
			_, ok := temp1[v]
			// 如果不存在直接返回false
			if !ok {
				return false
			} else { // 如果字符出现的次数超过一次,进行减1
				if temp1[v] > 0 {
					temp1[v]--
				} else {
					// 否则如果出现的次数已经不足一个证明,出现的次数不对,返回false
					return false
				}
			}
		}
		// 最后返回true即可
		return true

	}
}
