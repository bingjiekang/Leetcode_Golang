package main

// 给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。

// 注意:
// 十六进制中所有字母(a-f)都必须是小写。
// 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
// 给定的数确保在32位有符号整数范围内。
// 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。

// 针对正负数进行不同方式的提取,通过提取求余的结果,来转化对应的数
func toHex(num int) string {
	// 用来存储16进制对应的数
	var base []rune = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	// 用来存储转换好的结果
	var sult []rune = make([]rune, 0)
	// 赋值temp不影响原来传入数据
	temp := num
	// 当数据大于0时,通过反复求余得到对应的16进制位
	if temp >= 0 {
		for temp/16 != 0 {
			sult = append(sult, base[temp%16])
			temp /= 16
		}
		// 最后一次不足16后是对本身的16进制转换
		sult = append(sult, base[temp])
		length := len(sult)
		// 最后将数组翻转就是对应的16进制数
		for i, j := 0, length-1; i <= j; {
			sult[i], sult[j] = sult[j], sult[i]
			i++
			j--
		}
	} else { // 当数小于0时
		// 对数据赋初值
		sult = []rune{'f', 'f', 'f', 'f', 'f', 'f', 'f', 'f'}
		// 对负数取正,-1用来控制寻找base的位置
		temp = (-1)*temp - 1
		// 用来确定更改对应sult的下标
		var tmi int = len(sult) - 1
		// 通过求余取余得到对应的16进制位
		for ; temp/16 != 0; tmi-- {
			sult[tmi] = base[15-(temp%16)]
			temp /= 16
		}
		// 得到最后不足16对应的16进制的结果
		sult[tmi] = base[15-temp]
	}
	// 返回字符串即可
	return string(sult)

}
