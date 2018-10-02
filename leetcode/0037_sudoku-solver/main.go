package main

//题目描述:https://leetcode-cn.com/problems/sudoku-solver/description/
func solveSudoku(board [][]byte) {
	if !fill(board, '1', 0) {
		panic("此题无解")
	}
}

//fill	往编号为block的block中填数字n
//	board:数独数据  n:空格中待填充的数字  block:数字不重复的3×3小块
func fill(board [][]byte, n byte, block int) bool {
	if block == 9 {
		// 所有的block都已经填满，成功找到了解
		return true
	}

	//'9'+1 即ASCII中数字9的下一位
	if n == '9'+1 {
		// block 已经被填满了，去填写 block+1
		return fill(board, '1', block+1)
	}

	rowBegin := (block / 3) * 3
	colBegin := (block % 3) * 3

	//检测block中,n是否存在
	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if board[r][c] == n {
				// block 中已经有 n 了，无需填写
				// 去填写 n+1
				return fill(board, n+1, block)
			}
		}
	}

	//block不存在n

	//遍历block寻找当前block能存放n的位置
	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			//(r,c)填充n,符合基本条件
			if isAvaliable(board, n, r, c) {
				//假设(r,c)填n是正解
				board[r][c] = n

				//递归填充n+1
				if fill(board, n+1, block) {
					//数独解成功
					return true
				} else {
					//(r,c)填n不是正解
					// 把 (r,c) 还原，以便以后把 n 移入下一个可行的位置
					board[r][c] = '.'
				}

			}
		}
	}

	// n 在此 block 中无处可放。
	// 返回 false ，让 n-1 调整位置。
	return false
}

//isAvaliable (r,c)位置能否存放n
func isAvaliable(board [][]byte, n byte, r, c int) bool {
	// 当前位置上已有1-9的字符
	if board[r][c] != '.' {
		return false
	}

	// (r,c) 所在的行和列不能有 n
	// 在这里就可以体现，挨个往block中填写的优势了。
	for i := 0; i < 9; i++ {
		if board[r][i] == n || board[i][c] == n {
			return false
		}
	}

	return true
}

func main() {

}
