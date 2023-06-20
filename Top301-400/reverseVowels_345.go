package main

import (
	"strings"
)

// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。

// 解析:双指针left从前往后,right从后向前,直到他们指向aeiou或AEIOU1里的字母则停下来,等待对方停止后交换即可

func reverseVowels(s string) string {
	// 将字符串转为切片
	tmp := []rune(s)
	// 获取长度
	length := len(tmp)
	// 双指针,指向头和尾
	left, right := 0, length-1
	// 开始查找
	for left < right {
		// 右指针从后往前找到第一个元音字母后停下来,或者找不到则right=left
		for right > left && !strings.Contains("aeiouAEIOU", string(tmp[right])) {
			right--
		}
		// 左指针从前往后找到第一个元音字母停下来,或者找不到则right=left
		for left < right && !strings.Contains("aeiouAEIOU", string(tmp[left])) {
			left++
		}
		// 如果此时left<right则证明找到,交换即可,同时再对left,right指针再进行一次移动
		if left < right {
			tmp[left], tmp[right] = tmp[right], tmp[left]
			left++
			right--
		} else { // 没找到说明已经没有了,则退出即可
			break
		}
	}
	// 返回字符串
	return string(tmp)

}
