package main

// 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
// 解析:定义1个hash表,存储每个字母出现的最后的下标,再次便利原字符串,如果出现了与对应值下标相同的下标,且之前没有出现过则返回,若不同(代表出现过)赋值-1

func firstUniqChar(s string) int {
	// 定义一个hash表
	var lt map[rune]int
	// 获取长度,申请空间
	length := len(s)
	lt = make(map[rune]int, length)
	// 更新最后字母的下标
	for i, v := range s {
		lt[v] = i
	}
	// 遍历字符串,如果是第一个则返回,不是则赋值-1
	for i, v := range s {
		if lt[v] == i && lt[v] != -1 {
			return i
		}
		lt[v] = -1
	}
	return -1
}
