package main

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

// 解析:获得长度,当左长度小于右长度时,使用双指针,对如果相同则同时加1,不同则右指针右移,直到两指针有一个结束,如果左指针结束则返回true,如果右指针结束了但左指针没结束,返回false,
func isSubsequence(s string, t string) bool {
	// 得到两个字符串的长度
	slength, tlength := len(s), len(t)
	// s的长度大于t时,不会满足,返回false
	if slength > tlength {
		return false
	} else if slength == tlength { // 如果相同,则比较两个是否相同,如不同则返回false,若相同则返回true
		if s == t {
			return true
		} else {
			return false
		}
	} else { // 当s的长度小于t时
		var ls, lt int = 0, 0
		for ls < slength && lt < tlength {
			if s[ls] == t[lt] { // 如果对应字符相同则同时加1
				ls++
				lt++
			} else { // 否则右指针往后遍历,左指针不动
				lt++
			}
		}
		return ls == slength // 最后如果ls左指针遍历到最后的节点,则返回true否则返回false
	}
}
