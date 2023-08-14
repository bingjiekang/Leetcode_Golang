package main

// 设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。

// 以下是井字游戏的规则：

// 玩家轮流将字符放入空位（" "）中。
// 第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
// "X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
// 当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
// 当所有位置非空时，也算为游戏结束。
// 如果游戏结束，玩家不允许再放置字符。
// 如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。

// 示例 1：

// 输入： board = ["O X"," XO","X O"]
// 输出： "X"

// 解析：NxN 当某一行或者某一列或者对角线都相同时，返回胜利方，当没有胜利方且棋盘没空格则返回平局Draw，如果没有获胜方且有空格证明棋盘没有结束，返回Pending。通过记录每一行、每一列和对角线对应字符出现的次数，如果某一字符出现的次数为N则返回这个字符，在遍历时同时记录空白字符的个数，如果没有长度为N的字符，则看是否有空白字符个数，如果有则返回Pending继续棋盘，如果没有空白字符则证明平局，返回Draw
func tictactoe(board []string) string {
	// hp记录每一行对应字符出现次数，lp记录每一列
	hp, lp := map[byte]int{}, map[byte]int{}
	// dhp记录左对角线字符出现次数，dlp记录右对角线
	dhp, dlp := map[byte]int{}, map[byte]int{}
	// 定义blank记录空白字符的数量
	var blank int
	// 记录棋盘长度
	length := len(board)
	// 双层循环根据对角线遍历n行n列，并同时记录空白的数量
	for i := 0; i < length; i++ {
		// 行列
		for j := 0; j < length; j++ {
			// 计数每一行字符出现的次数
			if board[i][j] == 'X' {
				hp['X']++
			} else if board[i][j] == 'O' {
				hp['O']++
			} else { // 记录空白字符出现的次数
				blank++
			}
			// 计数每一列字符出现的次数
			if board[j][i] == 'X' {
				lp['X']++
			} else if board[j][i] == 'O' {
				lp['O']++
			}
		}
		// 如果这一行或者这一列出现的次数等于长度Length，返回这个字符
		if hp['X'] == length || lp['X'] == length {
			return "X"
		} else if hp['O'] == length || lp['O'] == length {
			return "O"
		} else { // 否则更新这些数为0
			hp['X'], lp['X'] = 0, 0
			hp['O'], lp['O'] = 0, 0
		}
		// 对角线
		// 左对角线，记录字符出现次数
		if board[i][i] == 'X' {
			dhp['X']++
		} else if board[i][i] == 'O' {
			dhp['O']++
		}
		// 右对角线，记录字符出现次数
		if board[i][length-1-i] == 'X' {
			dlp['X']++
		} else if board[i][length-1-i] == 'O' {
			dlp['O']++
		}
	}
	// 如果对角线出现字符次数等于length，返回这个字符
	if dhp['X'] == length || dlp['X'] == length {
		return "X"
	} else if dhp['O'] == length || dlp['O'] == length {
		return "O"
	} else if blank == 0 { // 否则就看空白字符出现的次数，若为0则平局
		return "Draw"
	} else { // 否则就是棋盘未结束，返回Pending
		return "Pending"
	}
}
