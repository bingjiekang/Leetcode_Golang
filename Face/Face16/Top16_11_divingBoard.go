package main

// 你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。
// 返回的长度需要从小到大排列。

// 输入：
// shorter = 1
// longer = 2
// k = 3
// 输出： [3,4,5,6]
// 解释：
// 可以使用 3 次 shorter，得到结果 3；使用 2 次 shorter 和 1 次 longer，得到结果 4 。以此类推，得到最终结果。

// 解析:判断k是否为0,若为0返回空,如果不为0,则判断shorter和longer,如果shorter==longer 则长度只有一种则总长度也就只有一种k*木板长度,如果shorter小于longer,则可以从0开始遍历到k,由于longer长木板出现的次数越多,木板越长,则遍历的时候总长度=longer*i+(k-i)*shorter,遍历得到所有的组合
func divingBoard(shorter int, longer int, k int) (sult []int) {
	// 如果k为0则直接返回
	if k == 0 {
		return
	}
	// 如果shorter和longer相同,则返回唯一结果
	if shorter == longer {
		sult = append(sult, shorter*k)
	} else if shorter < longer { // 否则的话则遍历,将对应数据加入即可,
		for i := 0; i <= k; i++ {
			sult = append(sult, longer*i+(k-i)*shorter)
		}
	}
	// 最后返回结果
	return
}
