package algorithm

import "fmt"

//回溯算法处理 8 皇后问题
var Result [8]int //全局变量，下标为行，值为列

func Cal8queens(row int) { //调用方式 Cal8queens(0)
	if row == 8 { //8个棋子都放好了，打印结果
		fmt.Println(Result)
		return
	}

	for column := 0; column < 8; column++ {
		//每一行都有 8 种放法
		if isOk(row, column) {
			Result[row] = column
			Cal8queens(row + 1)
		}
	}
}

//判断 row 行 column 列是否合法
func isOk(row int, column int) bool {
	leftUp := column - 1
	rightUp := column + 1

	//逐行考察每一行
	for i := row - 1; i >= 0; i-- {
		if Result[i] == column {
			//第 i 行的 column 列有棋子，不满足竖线
			return false
		}
		if leftUp >= 0 {
			if Result[i] == leftUp {
				//对角线有棋子，不满足
				return false
			}
		}
		if rightUp < 8 {
			if Result[i] == rightUp {
				//对角线有棋子，不满足
				return false
			}
		}
		leftUp -= 1
		rightUp += 1
	}
	return true
}
