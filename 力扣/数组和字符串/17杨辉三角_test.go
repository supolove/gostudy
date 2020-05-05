package 数组和字符串

import (
	"testing"
)

/**
https://leetcode-cn.com/explore/learn/card/array-and-string/202/conclusion/792/
*/

func getRow(rowIndex int) []int {

	var temArr []int = nil
	for i := 0; i < rowIndex+1; i++ {
		arr := make([]int, i+1)
		for j := 0; j < len(arr); j++ {
			if temArr == nil {
				arr[j] = 1
			} else {
				/*
					switch j {
					case 0:
						arr[j] = temArr[j]
					case len(temArr):
						arr[j] = temArr[j-1]
					default:
						arr[j] = temArr[j-1] + temArr[j]
					}
				*/
				if j == 0 {
					arr[j] = temArr[j]
				} else if j == len(temArr) {
					arr[j] = temArr[j-1]
				} else {
					arr[j] = temArr[j-1] + temArr[j]
				}
			}
		}
		//fmt.Println(arr)
		temArr = arr
	}
	return temArr
}

func Test_getRow(t *testing.T) {
	getRow(3)
}
