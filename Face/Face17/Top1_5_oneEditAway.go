package main

// 字符串有三种编辑操作:插入一个英文字符、删除一个英文字符或者替换一个英文字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

// 示例 1:
// 输入:
// first = "pale"
// second = "ple"
// 输出: True

// 解析:先分析得出三种情况
// 第一种:两个长度差大于1,返回false
// 第二种:两个长度差为0,有两种情况,若字符串相同返回true,若不同则计算不同的个数,若不同的个数大于1则返回false [编辑:修改]
// 第三种:两个长度差为1,则用双指针分别指向first和second,当出现不同时,谁的长谁移动,并记录不同的次数,当大于1时返回false [编辑:增加或者删除]
// 最后都不满足时返回true

func oneEditAway(first string, second string) bool {
	// 需要一个计数用来查看修改的次数，如果大于1次返回false
	var cout int
	// 获取两个字符串的长度并记录
	length1, length2 := len(first), len(second)
	record := abs(length1, length2)
	// 差值大于1个返回false
	if record > 1 {
		return false
	} else if record == 0 {
		// 差值为0，则是替换或相同
		if first == second { // 字符串相同
			return true
		} else { // 不同需要替换
			for i := 0; i < length1; i++ {
				// 对应字符不想同,计数cout
				if first[i] != second[i] {
					cout++
					// cout大于一次返回false
					if cout > 1 {
						return false
					}
				}

			}
		}
	} else { // 差值为1，看哪一个长，哪一个长，指针移动的时候移动谁
		// 双指针遍历first和second
		for i, j := 0, 0; i < length1 && j < length2; {
			// 如果对应字符不同，谁长移动谁
			if first[i] != second[j] {
				if length1 > length2 {
					i++
				} else {
					j++
				}
				cout++
				// cout大于一返回false
				if cout > 1 {
					return false
				}
			} else { // 相同则一起移动
				i++
				j++
			}

		}

	}
	return true
}

// 用来获取两个长度的差值
func abs(n, m int) int {
	if n >= m {
		return n - m
	}
	return m - n
}
