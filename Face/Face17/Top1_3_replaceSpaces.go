package face17

// URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

// 示例 1：

// 输入："Mr John Smith    ", 13
// 输出："Mr%20John%20Smith"
// 示例 2：

// 输入："               ", 5
// 输出："%20%20%20%20%20"

// 解析:经过分析得出,出现空格则用%20替换 最终的替换后的总数为指定数量,那么解决方法就是从前往后,检查到空格就用%20加入到答案切片中,其他字符就直接加入,当加入的数量等于指定数量时,直接停止返回即可.

func replaceSpaces(S string, length int) string {
	// 用来计数,判断已经转换的字符数
	var cout int
	// 用来存储新转换的数
	var sult []rune = make([]rune, 0)
	// 遍历字符串
	for _, v := range S {
		// 如果是空格则加入%20到sult中
		if v == ' ' {
			sult = append(sult, '%')
			sult = append(sult, '2')
			sult = append(sult, '0')

		} else { // 否则就直接加入
			sult = append(sult, v)
		}
		cout++
		// 如果次数和指定长度一样,返回即可
		if cout == length {
			return string(sult)
		}
	}
	return ""
}
