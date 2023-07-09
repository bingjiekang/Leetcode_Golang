package main

// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
// 如果可以，返回 true ；否则返回 false 。
// magazine 中的每个字符只能在 ransomNote 中使用一次。

// 解析:利用hash表存储并解析,先对magazine进行遍历存储,然后对ransomNote进行遍历,每出现一个减1,直到为0或者不存在,返回false,若全部遍历后,则返回true
func canConstruct(ransomNote string, magazine string) bool {
	// 定义一个hash表
	var hash map[rune]int = make(map[rune]int, 0)
	// 遍历magazine并存储
	for _, v := range magazine {
		hash[v]++
	}
	// 遍历ransomNote对出现的进行减1判断
	for _, v := range ransomNote {
		// _,err := hash[v]
		// 如果存在则减1
		if hash[v] > 0 {
			hash[v]--
		} else { // 不存在或者已经等于0,则直接返回false
			return false
		}
	}
	return true
}
