package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/257/recurrence-relation/1203/
*/

func digui2(numRows int, upRow []int) []int {
	if numRows > 1 {
		upRow = digui2(numRows-1, upRow)
	}
	curRow := []int{}
	if numRows == 1 {
		curRow = append(curRow, 1)
	} else {
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
	return curRow
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := []int{}
	result = digui2(rowIndex+1, result)
	return result
}

func Test_getRow(t *testing.T) {
	fmt.Println(getRow(4))
}
