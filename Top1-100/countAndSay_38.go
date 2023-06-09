package main

// 一个正整数 n ，输出外观数列的第 n 项。
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。
// 你可以将其视作是由递归公式定义的数字字符串序列：
// countAndSay(1) = "1"
// countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 第一项是数字 1
// 描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
// 描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
// 描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
// 描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
// 解析:对于每一层,用两个数组分别存储某个数出现几次和这个数是几,每一层结束时,将对应数据加入到sult结果中,然后反复更新,找到需要的第n层
func countAndSay(n int) string {
	// tml用来存储数字出现的次数
	tml := []byte{'1'}
	// tmr用来存储数字为几
	tmr := []byte{'1'}
	// sult用来存储每一层的结果
	var sult []rune = []rune{'1'}
	// 用来确定每一层需要遍历的次数
	var length int
	// 最外层,用来获得第n个外观数字
	for i := 1; i < n; i++ {
		// 获取每层的长度
		length = len(sult)
		// 用来确定数字出现的次数
		temp := 1
		// 对于该层比价当前数和前一个数
		for li := 1; li < length; li++ {
			// 如果相等将 temp次数加1
			if sult[li] == sult[li-1] {
				temp++
			} else { // 不等的话 就将前一个数字出现的次数和前一个数字,加入到sult中,已unicode数字对应的码数加入
				tml = append(tml, byte(temp)+48)
				tmr = append(tmr, byte(sult[li-1]))
				// 同时重置temp次数为1
				temp = 1
			}
			// 如果当前遍历到最后一位,则直接将本位加入
			if li == length-1 {
				tml = append(tml, byte(temp)+48)
				tmr = append(tmr, byte(sult[li]))
			}
		}
		// 用来对每一层,数字出现的次数和数字加入到sult中
		for i, v := range tml {
			sult = append(sult, rune(v))
			sult = append(sult, rune(tmr[i]))
		}
		// 更新每一层sult的值,将此前的值去掉
		sult = sult[length:]
		// 更新tml和tml值用来存储新的值
		tml = []byte{}
		tmr = []byte{}

	}
	// 将rune转为字符串返回即可
	return string(sult)

}
