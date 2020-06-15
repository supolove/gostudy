package _0_06

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode-cn.com/problems/longest-common-prefix/
*/

func longestCommonPrefix(strs []string) string {

	if strs == nil || len(strs) == 0 {
		return ""
	}
	builder := strings.Builder{}

flag:
	for i := 0; i < len(strs[0]); i++ {
		var temp byte
		for j := 0; j < len(strs); j++ {
			if i < len(strs[j]) {
				if temp == 0 {
					temp = strs[j][i]
				} else if temp != strs[j][i] {
					break flag
				}
			} else {
				break flag
			}
		}
		builder.WriteByte(strs[0][i])
	}
	return builder.String()
}

func Test_longestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix(nil))
}
