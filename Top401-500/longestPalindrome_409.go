package top401500

// 给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。
// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。

// 输入:s = "abccccdd"
// 输出:7
// 解释:
// 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

// 解析:用hash表存储每个字符出现的次数,然后遍历这个hash表,如果出现的次数是2的倍数,证明是回文数直接加入他的长度,如果不是2的倍数,那么对他进行减1操作,那么此时他要么为0,要么为偶数,则是回文数,加入即可,然后记录此时有个多余的数 为1,最后加入即可

func longestPalindrome(s string) int {
	// 定义一个hash表,用来存储字符出现的次数
	var hash map[rune]int = make(map[rune]int, 0)
	// 定义连个变量,一个存储回文字符串的次数,一个存储不是回文串的中间的那个数1
	var sult, temp int
	// 遍历字符串,将字符串出现的次数存到hash表中
	for _, v := range s {
		hash[v]++
	}
	// 遍历这个hash表
	for i := range hash {
		// 如果是2的倍数,直接加入
		if hash[i]%2 == 0 {
			sult += hash[i]
		} else { // 不是2的倍数,则减1后加入
			sult += (hash[i] - 1)
			temp = 1 // 将中间的那个数记录下来保持为1
		}
	}
	// 返回结果即可
	return sult + temp

}
