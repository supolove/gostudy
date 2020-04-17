package 数组和字符串

import (
	"fmt"
	"testing"
)

/**

  最长公共前缀

https://leetcode-cn.com/explore/learn/card/array-and-string/200/introduction-to-string/781/
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; ; i++ {
		var a byte = 0
		for j := 0; j < len(strs); j++ {
			if i < len(strs[j]) {
				if a == 0 {
					a = strs[j][i]
				} else if a != strs[j][i] {
					return strs[0][:i]
				}
			} else {
				return strs[0][:i]
			}
		}
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	//str := []string{"ooower","flow","flight"}
	str := []string{"a", "ab"}
	fmt.Println(longestCommonPrefix(str))
}
