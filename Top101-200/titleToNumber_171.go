package main

/*
要求:给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。
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

func titleToNumber(columnTitle string) int {
	// 可以看做是26进制转为10进制 num = 26^(n-1)+26^(n-2)....+26^0
	length := len(columnTitle) - 1
	var sums int
	for i := length; i >= 0; i-- {
		temp := 1
		// 用来实现26的n次方
		for tm := length - i; tm > 0; tm-- {
			temp *= 26
		}
		// 获得26进制转换的10进制数
		sums += int(columnTitle[i]-'A'+1) * temp // 对‘A’字符进行加减后记得int强制转换一下
	}
	return sums
}
