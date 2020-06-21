package _0_06

import (
	"fmt"
	"testing"
)

/*

https://leetcode-cn.com/problems/regular-expression-matching/
 正则表达式
 . 表示任意一个字符
 * 表示零个或多个前面的那一个元素
 .* 贪婪匹配，可以先转化为任意多个 点，相当于可以匹配任意字符串
*/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func Test_isMatch(t *testing.T) {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
