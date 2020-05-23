package 哈希表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/206/practical-application-design-the-key/822/
*/

func isValid(b []byte) bool {
	m := map[byte]bool{}
	for _, bb := range b {
		if bb == '.' {
			continue
		}
		_, ok := m[bb]
		if ok {
			return false
		}
		m[bb] = true
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	// 横
	for _, a := range board {
		if !isValid(a) {
			return false
		}
	}

	// 竖
	var column [][]byte
	for i := 0; i < 9; i++ {
		var a []byte
		for _, as := range board {
			a = append(a, as[i])
		}
		if !isValid(a) {
			return false
		}
		column = append(column, a)
	}

	// 方
	for i := 0; i < 9; i++ {
		var a []byte
		if 0 <= i && i < 3 {
			for j := 0; j < 3; j++ {
				a = append(a, board[j][3*i:3*(i+1)]...)
			}
			if !isValid(a) {
				return false
			}
		} else if 3 <= i && i < 6 {
			for j := 0; j < 3; j++ {
				a = append(a, board[j+3][3*(i-3):3*((i-3)+1)]...)
			}
			if !isValid(a) {
				return false
			}
		} else {
			for j := 0; j < 3; j++ {
				a = append(a, board[j+6][3*(i-6):3*((i-6)+1)]...)
			}
			if !isValid(a) {
				return false
			}
		}

	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	data := [3][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			val := 1 << uint(board[i][j]-'0')
			if data[0][i]&val == val || data[1][j]&val == val || data[2][i/3*3+j/3]&val == val {
				return false
			}

			data[0][i] = data[0][i] | val
			data[1][j] = data[1][j] | val
			data[2][i/3*3+j/3] = data[2][i/3*3+j/3] | val
		}
	}
	return true
}

func Test_isValidSudoku(t *testing.T) {
	arr := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(arr))
}
