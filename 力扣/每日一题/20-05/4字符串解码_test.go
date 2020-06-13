package _0_05

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
2020-05-28
https://leetcode-cn.com/problems/decode-string/
*/

func decodeString(s string) string {
	result := s
	num := 0
	for strings.Contains(result, "[") {

		// 获取数字项
		start := strings.IndexByte(result, '[')
		for i := 0; i <= start; i++ {
			if result[i] >= '0' && result[i] <= '9' {
				num, _ = strconv.Atoi(result[i:start])
				result = result[:i] + result[start:]
				break
			}
		}

		start = strings.IndexByte(result, '[')
		record := 0

		for i := start + 1; i < len(result); i++ {
			if result[i] == '[' {
				record--
			}

			if result[i] == ']' {
				// 找到匹配
				if record == 0 {
					// 拼接
					v := result[start+1 : i]
					result = result[:start] + strings.Repeat(v, num) + result[i+1:]
					break
				} else {
					record++
				}
			}
		}
	}
	return result
}

func Test_decodeString(t *testing.T) {
	fmt.Println(decodeString("3[a2[c]]"))
}
