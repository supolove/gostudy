package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
二进制求和
https://leetcode-cn.com/explore/learn/card/array-and-string/200/introduction-to-string/779/
*/

func addBinary(a string, b string) string {

	temp := make([]byte, 0)
	for i := 0; ; i++ {

		var aa byte = 0
		if i < len(a) {
			aa = a[len(a)-i-1] - 48
		}

		var bb byte = 0
		if i < len(b) {
			bb = b[len(b)-i-1] - 48
		}

		temp = append(temp, aa+bb)
		if i >= len(b)-1 && i >= len(a)-1 {
			break
		}
	}

	result := make([]byte, len(temp)+1)
	for i := 0; i < len(temp); i++ {
		v := temp[i]
		if v+result[len(result)-i-1] == 2 {
			result[len(result)-i-1] = 0
			result[len(result)-i-2]++
		} else if v+result[len(result)-i-1] == 3 {
			result[len(result)-i-1] = 1
			result[len(result)-i-2]++
		} else {
			result[len(result)-i-1] += v
		}
	}
	if result[0] == 0 {
		result = result[1:]
	}

	for i := 0; i < len(result); i++ {
		result[i] += 48
	}

	return string(result)
}

func Test_addBinary(t *testing.T) {
	fmt.Println(addBinary("1010", "1011"))
}
