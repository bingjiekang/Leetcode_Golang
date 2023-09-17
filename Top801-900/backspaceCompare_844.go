package main

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

// 注意：如果对空文本输入退格字符，文本继续为空。
// 示例 1：
// 输入：s = "ab#c", t = "ad#c"
// 输出：true
// 解释：s 和 t 都会变成 "ac"。

// 模拟栈，得到对应结果，出现#，删除前面一个字母
// func backspaceCompare(s string, t string) bool {
// 	var lts, ltt []rune = make([]rune, 0), make([]rune, 0)
// 	var couts, coutt int
// 	for _, v := range s {
// 		if v == '#' && couts > 0 {
// 			lts = lts[0 : couts-1]
// 			couts--
// 		} else if v != '#' {
// 			lts = append(lts, v)
// 			couts++
// 		}
// 	}

// 	for _, v := range t {
// 		if v == '#' && coutt > 0 {
// 			ltt = ltt[0 : coutt-1]
// 			coutt--
// 		} else if v != '#' {
// 			ltt = append(ltt, v)
// 			coutt++
// 		}
// 	}
// 	if string(lts) == string(ltt) {
// 		return true
// 	} else {
// 		return false
// 	}

// }

// 双指针，一个字符是否会被删掉，只取决于该字符后面的退格符，而与该字符前面的退格符无关。因此当我们逆序地遍历字符串，就可以立即确定当前字符是否会被删掉。我们定义两个指针，分别指向两字符串的末尾。每次我们让两指针逆序地遍历两字符串，直到两字符串能够各自确定一个字符，然后将这两个字符进行比较。重复这一过程直到找到的两个字符不相等，或遍历完字符串为止。

func backspaceCompare(s string, t string) bool {
	lengths, lengtht := len(s)-1, len(t)-1
	rgs, rgt := 0, 0
	for lengths >= 0 || lengtht >= 0 {
		// 找到对应的删除后的第一个值，比较
		for lengths >= 0 {
			if s[lengths] == '#' {
				rgs++
				lengths--
			} else {
				if rgs != 0 {
					rgs--
					lengths--
				} else {
					break
				}
			}

		}

		for lengtht >= 0 {
			if t[lengtht] == '#' {
				rgt++
				lengtht--
			} else {
				if rgt != 0 {
					rgt--
					lengtht--
				} else {
					break
				}
			}

		}
		// 如果都大于等于0，且对应值不同，返回false
		if lengths >= 0 && lengtht >= 0 && s[lengths] != t[lengtht] {
			return false
		} else if lengths >= 0 && lengtht >= 0 && s[lengths] == t[lengtht] { // 相同的话就更新下标
			lengths--
			lengtht--
		} else if lengths >= 0 || lengtht >= 0 { // 否则 就是一个小于0，一个大于等于0，返回false
			return false
		}

	}
	return true

}
