package main

/*
要求:
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。字母和数字都属于字母数字字符。
*/

// 判断是否是回文字符串
func isPalindrome(s string) bool {
	str := tranfrom(s)
	length := len(str) - 1
	for i := 0; i <= length; i++ {
		if str[i] != str[length] {
			return false
		}
		length--
	}
	return true
}

// 将字符串转为只有小写和数字的字符串
func tranfrom(s string) string {
	str := []rune(s)
	var str1 []rune
	for _, v := range str {
		// 大写字母转为小写字母加入
		if v >= 65 && v <= 90 {
			str1 = append(str1, v+32)
			// 小写字母和数字直接加入
		} else if (v >= 48 && v <= 57) || (v >= 97 && v <= 122) {
			str1 = append(str1, v)
		}
	}
	return string(str1)
}
