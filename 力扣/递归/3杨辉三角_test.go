package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/257/recurrence-relation/1202/
*/

func digui(numRows int, table [][]int) [][]int {
	if numRows > 1 {
		table = digui(numRows-1, table)
	}
	curRow := []int{}
	if numRows == 1 {
		curRow = append(curRow, 1)
	} else {
		upRow := table[numRows-2]
		for i := 0; i < numRows; i++ {
			if i == 0 {
				curRow = append(curRow, upRow[i])
			} else if i == numRows-1 {
				curRow = append(curRow, upRow[i-1])
			} else {
				curRow = append(curRow, upRow[i-1]+upRow[i])
			}
		}
	}
	table = append(table, curRow)

	return table
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := [][]int{}
	result = digui(numRows, result)
	return result
}

func Test_yhsj(t *testing.T) {
	result := generate(0)
	fmt.Println(result)
}
