package main

// 珠玑妙算游戏（the game of master mind）的玩法如下。
// 计算机有4个槽，每个槽放一个球，颜色可能是红色（R）、黄色（Y）、绿色（G）或蓝色（B）。例如，计算机可能有RGGB 4种（槽1为红色，槽2、3为绿色，槽4为蓝色）。作为用户，你试图猜出颜色组合。打个比方，你可能会猜YRGB。要是猜对某个槽的颜色，则算一次“猜中”；要是只猜对颜色但槽位猜错了，则算一次“伪猜中”。注意，“猜中”不能算入“伪猜中”。

// 给定一种颜色组合solution和一个猜测guess，编写一个方法，返回猜中和伪猜中的次数answer，其中answer[0]为猜中的次数，answer[1]为伪猜中的次数。

// 示例：

// 输入： solution="RGBY",guess="GGRR"
// 输出： [1,1]
// 解释： 猜中1次，伪猜中1次。

// len(solution) = len(guess) = 4
// solution和guess仅包含"R","G","B","Y"这4种字符

// 解析:通过字典,遍历一次,如果对应相同则猜中加1次,如果对应不同,则判断相应字符是否在另一个字典中出现,如果出现,则将对应字符出现次数减1,将本字典中对应字符出现次数减1,相应的伪猜中次数加1即可

func masterMind(solution string, guess string) []int {
	// solu存储实际的,gues存储猜测的
	var solu, gues map[byte]int = make(map[byte]int, 0), make(map[byte]int, 0)
	// gues_t存储猜中的次数,gues_f存储伪猜中的次数
	var coutgues_t, coutgues_f int
	for i := 0; i < 4; i++ {
		// 相应字典加1
		solu[solution[i]]++
		gues[guess[i]]++
		// 如果相同则操作猜中次数加1,对应字符减1
		if solution[i] == guess[i] {
			coutgues_t++
			solu[solution[i]]--
			gues[guess[i]]--
		} else { // 如果不同,则判断自己的字符是否对应另外字典中出现,若出现则增加伪猜中次数,若不出现则不改变,因为对应字符不同,所以不用双重判断对方是否在自己字典中
			if solu[guess[i]] > 0 {
				coutgues_f++
				solu[guess[i]]--
				gues[guess[i]]--
			}
			if gues[solution[i]] > 0 {
				coutgues_f++
				solu[solution[i]]--
				gues[solution[i]]--
			}
		}
	}
	return []int{coutgues_t, coutgues_f}
}
