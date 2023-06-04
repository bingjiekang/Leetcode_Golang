package main

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 解析:通过hash表进行比较,可以通过简易的列表,对每个小写字母进行-'a'进行数字化处理,对应列表里的数据,最后比较两个列表是否相同得到
func isAnagram(s string, t string) bool {
	// 两个列表用来存储两个字符串的字符出现的次数
	var lf, rh [26]int
	// 分别记录
	for _, v := range s {
		lf[v-'a']++
	}
	for _, v := range t {
		rh[v-'a']++
	}
	// 如果两个列表相同返回true否则返回false
	return lf == rh

}
