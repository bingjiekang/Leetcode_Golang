package face10

// 面试题 10.05. 稀疏数组搜索
// type: 简单

// 稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
// 示例1:
//  输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
//  输出：-1
//  说明: 不存在返回-1。
// 示例2:

//  输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
//  输出：4

// 使用键值对形式存储对应字符串出现的下标,如果为空字符串"",则跳过,最后查询看是否有这个字符串,如果有就返回值,没有就返回-1
func findString(words []string, s string) int {
	var dict map[string]int = make(map[string]int, 0)
	for k, v := range words {
		if v == "" {
			continue
		}
		dict[v] = k
	}
	_, ok := dict[s]
	if ok {
		return dict[s]
	}
	return -1
}
