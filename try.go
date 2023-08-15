package main

import "fmt"

func main() {
	fmt.Println(9/5)
}

func tictactoe(board []string) string {
	// 定义blank记录空白字符的数量
	var blank int
	// 双层循环根据对角线遍历三行三列，并同时记录空白的数量
	for i := 0; i < 3; {
		if (board[i][0] == board[i][1] && board[i][0] == board[i][2]) || (board[0][i] == board[1][i] && board[0][i] == board[2][i]) {
			return string(board[i][i])
		} else {
			for j := 0; j < 3; j++ {
				if board[i][j] == ' ' {
					blank++
				}
			}
		}
	}
	if (board[0][0] == board[1][1] && board[0][0] == board[2][2]) || (board[0][2] == board[1][1] && board[1][1] == board[2][0]) {
		return string(board[1][1])
	} else {
		if blank == 0 {
			return "Draw"
		} else {
			return "Pending"
		}
	}

	// 在记录对角线是否合理，若合理某一方获胜

	// 如果blank为0则返回Pending，若不为0切不获胜返回返回Draw

}
