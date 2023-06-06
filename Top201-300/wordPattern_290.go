package main

import (
	"strings"
)

// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
// 示例
// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true

// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false
// 利用俩个哈希表,将左右两边分别对应key和value,存储到对应的hash表中,都不存在就加入,有任何一个存在就比较对应的hash表a对应的值和对应的另一个hash表b对应的的key是否相同即可
func wordPattern(pattern string, s string) bool {
	// s是带空格的字符串,采用split函数分割存储到列表中,便于取值和比较
	var tmp []string = strings.Split(s, " ")
	// 初始化两个hash表
	var left map[rune]string = make(map[rune]string, 1)
	var right map[string]rune = make(map[string]rune, 1)
	// 长度不同则一定不会对应,返回false即可
	if len(pattern) != len(tmp) {
		return false
	}
	//遍历其中一个字符串
	for i, v := range pattern {
		// 获取对应的hash表是否存在当前key的值
		_, ok := left[v]
		_, okt := right[tmp[i]]
		// 都不存在的时候加入对应的value
		if !ok && !okt {
			left[v] = tmp[i]
			right[tmp[i]] = v
		} else { // 否则就比较对应的hash表a和hash表b对应的值是否一一对应,不对应返回false
			if left[v] != tmp[i] || right[tmp[i]] != v {
				return false
			}
		}
	}
	// 全部符合 返回true
	return true
}
