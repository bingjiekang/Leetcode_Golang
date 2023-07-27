package main

import (
	"fmt"
	"math/bits"
)

// 二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。

// 示例 1：
// 输入：turnedOn = 1
// 输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]

// 枚举：外层遍历小时，内层遍历分钟，将对应的数字得到对应的2进制位数与对应给的2进制位进行比较，若相同则加人字符串数组，我们可以枚举小时的所有可能值 [0,11][0,11][0,11]，以及分钟的所有可能值 [0,59][0,59][0,59]，并计算二者的二进制中 111 的个数之和，若为 turnedOn\textit{turnedOn}turnedOn，则将其加入到答案中。

func readBinaryWatch(turnedOn int) (sult []string) {
	// 定义i,j为无符号32位
	var i, j uint
	// 外层遍历小时
	for i = 0; i < 12; i++ {
		// 内层遍历分钟
		for j = 0; j < 60; j++ {
			// 函数OnesCount(需要为uint类型)得到对应数的2进制个数,如果相加后和turnedOn相同则加入到数组里
			if bits.OnesCount(i)+bits.OnesCount(j) == turnedOn {
				sult = append(sult, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return
}
