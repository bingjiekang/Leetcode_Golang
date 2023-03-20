package main

/*
要求:给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
例如：
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
*/

// 26个英文字母
const dow int = 26

func convertToTitle(columnNumber int) string {
	lt := transform(columnNumber)
	// 定义切片用来转换字符串
	var temp []byte = make([]byte, 0)
	length := len(lt) - 1
	// 将数组数据分别加入切片
	for i := length; i >= 0; i-- {
		temp = append(temp, 'A'+byte(lt[i]))
	}
	return string(temp)

}

func transform(num int) (list []int) {
	temp := num
	tm := -1
	// 用来获取每个数字
	for temp > 0 {
		// 假设A按照0来，所以对应的其他属都需要减一
		temp--
		tm = temp % dow
		temp = temp / dow
		list = append(list, tm)
	}
	return list
}
