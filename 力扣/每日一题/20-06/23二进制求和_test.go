package _0_06

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/add-binary/
*/

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	max := maxv(la, lb)
	res := make([]byte, max+1)
	for i := 0; i < len(res); i++ {
		res[i] = '0'
	}
	for i := len(res) - 1; i >= 0; i-- {

		var va byte

		last := len(res) - 1 - i

		if la-1-last >= 0 {
			va += a[la-1-last] - '0'
		}
		if lb-1-last >= 0 {
			va += b[lb-1-last] - '0'
		}
		va += res[i] - '0'

		if va == 3 {
			res[i] = '1'
			res[i-1] = '1'
		} else if va == 2 {
			res[i] = '0'
			res[i-1] = '1'
		} else if va == 1 {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}

	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}
func maxv(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_addBinary(t *testing.T) {
	fmt.Println(addBinary("1111", "1111"))
}
