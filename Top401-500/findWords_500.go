package top401500

import "strings"

// 500. 键盘行
// 简单
// 给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。

// 美式键盘 中：
// 第一行由字符 "qwertyuiop" 组成。
// 第二行由字符 "asdfghjkl" 组成。
// 第三行由字符 "zxcvbnm" 组成。
// 示例 1：
// 输入：words = ["Hello","Alaska","Dad","Peace"]
// 输出：["Alaska","Dad"]
// 先用map存储对应每一行键盘的值，第一行0，第二行1，第三行2，遍历对应数组，记录字符串里的字符，如果都相同就加入sult，如果不同就跳出来
var strand []string = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

func findWords(words []string) (sult []string) {
	var dict map[rune]int = make(map[rune]int, 0)
	// 记录对应的字母对应第几行
	for k, v := range strand {
		for _, t := range v {
			dict[t] = k
		}
	}
	// 遍历每个字符串，判断所处是否在一行
	for _, v := range words {
		// 转换对应大写的字符
		vt := strings.ToLower(v)
		target := dict[rune(vt[0])]
		success := true
		for _, tmp := range vt {
			// 如果不在一行就break
			if target != dict[tmp] {
				success = false
				break
			}
		}
		// 如果都相同 则加入即可
		if success {
			sult = append(sult, v)
		}
	}
	return sult
}
