package main

// 给定两个字符串 s 和 t ，它们只包含小写字母。
// 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
// 请找出在 t 中被添加的字母。

// 解析:用hash表存储数据1,对每个字符进行计数存储,然后遍历数据2,对出现的进行减1,如果出现不同的或者该字符在之前出现的次数已经符合,则返回这个字符

func findTheDifference(s string, t string) (sult byte) {
	// 定义hash表,存储字符串s的数据及字符出现的次数
	var hash map[rune]int = make(map[rune]int, 0)
	// 对s字符串进行遍历存储
	for _, v := range s {
		hash[v] += 1
	}
	// 对t进行遍历
	for _, v := range t {
		_, err := hash[v]
		// 如果对应的字符不存在或者在前面出现的次数已经足够则进行返回
		if !err || hash[v] == 0 {
			sult = byte(v)
			break
		} else { // 否则,对对应的字符进行计数减1
			hash[v]--
		}
	}
	return sult

}
