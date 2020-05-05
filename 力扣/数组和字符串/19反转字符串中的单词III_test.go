package 数组和字符串

import (
	"fmt"
	"strings"
	"testing"
)

func reverseWords3(s string) string {
	temp := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && i == 0 {
			continue
		}
		if s[i] == ' ' || i == len(s)-1 {
			var ss = ""
			if i == len(s)-1 {
				ss = s[start : i+1]
			} else {
				ss = s[start:i]
			}
			start = i + 1
			t := make([]byte, 0)
			for j := len(ss) - 1; j >= 0; j-- {
				t = append(t, ss[j])
			}
			temp = append(temp, string(t))
		}
	}
	return strings.Join(temp, " ")
}

func reverseWords4(s string) string {
	srclist := strings.Split(s, " ")
	for j, v := range srclist {
		tmp := []byte(v)
		for i := 0; i < len(tmp)/2; i++ {
			tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
		}
		srclist[j] = string(tmp)
	}
	return strings.Join(srclist, " ")
}

func Test_reverseWords3(t *testing.T) {
	fmt.Println(reverseWords3("Let's take LeetCode contest"))
}
