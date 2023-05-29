package main

// 给定两个字符串 s 和 t ，判断它们是否是同构的。
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
// 示例 1:
// 输入：s = "egg", t = "add"
// 输出：true
// 示例 2：
// 输入：s = "foo", t = "bar"
// 输出：false

// 解析:用两个哈希表来维护,由于长度相同,将对应的key和value分别为对应的字母,如果不存在则分别加入,如果存在,则判断已经存在的每个对应key的value和当前另一个对应的字母是否相同,不同的话,证明不是同构字符串

func isIsomorphic(s string, t string) bool {
	// 获取字符串的长度
	tmp := len(s)
	// 由于字符串是编码形式存储,则声明两个map[byte]byte{},byte类型的的字典存储信息
	tpleft, tpright := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < tmp; i++ {
		// 将对应的每个字母分别存储一下,便于后续引用
		x, y := s[i], t[i]
		// 如果当前字符已经存在,则判断当前字符(key)对应的value是否和另一个字符串的对应顺序字符相同,若不同则,证明一边重复出现的字符在另一边并没有重复出现
		if (tpleft[x] > 0 && tpleft[x] != y) || (tpright[y] > 0 && tpright[y] != x) {
			return false
		}
		// 如果当前字符不存在,则将当前的字符对应的另一个字符串的字符,分别作为key和value分别加入到两个hash表中
		tpleft[x] = y
		tpright[y] = x
	}
	return true
}
