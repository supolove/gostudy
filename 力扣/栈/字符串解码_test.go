package 栈

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 我的实现
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

// 递归实现
/*
func decodeString_digui(s string,index , repeat int) string  {
	res := ""
	next_repeat := 0
	for index < len(s)  {
		if s[index] == '[' {
			index ++
			res = res + decodeString_digui(s, index, next_repeat)
			next_repeat = 0
		} else if s[index] >= '0' && s[index] <= '9' {
			index ++
			c := s[index]
			next_repeat = next_repeat * 10 + (c - '0')
		}else if s[index] == ']'{
			index ++
			tmp := res
			for i:=1; i< repeat; i++{
				res = res + temp
			}
			return res
		}else{
			index++
			res = res + s[index]
		}
	}
	return res
}

*/

func Test_decodeString(t *testing.T) {
	s := "100[leetcode]"
	fmt.Println(decodeString(s))
}
