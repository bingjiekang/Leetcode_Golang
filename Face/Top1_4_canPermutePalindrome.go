package main

// 给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

// 回文串不一定是字典当中的单词。

// 示例1：

// 输入："tactcoa"
// 输出：true（排列有"tacocat"、"atcocta"，等等）

// 解析:通过map存储字符出现的次数,如果字符出现的次数为奇数,且这样的字符大于1个则返回false,否则全部符合则为true

func canPermutePalindrome(s string) bool {

	// 用一个变量记录是否是回文串，一个变量记录奇数次字符的次数
	var cout int
	// 使用map进行存储字符出现的次数
	var hash map[rune]int = make(map[rune]int, 0)
	for _, v := range s {
		hash[v]++
	}
	// 遍历map
	for _, v := range hash {
		// 如果都是偶数，则是回文排列
		if v%2 != 0 {
			cout++
			if cout > 1 {
				return false
			}
		}

	}
	return true

}
