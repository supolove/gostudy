package 数组和字符串

import (
	"strings"
	"testing"
)

func reverseWords(s string) string {
	arr := strings.Fields(s)
	temp := make([]string, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		if strings.TrimSpace(arr[i]) != "" {
			temp = append(temp, strings.TrimSpace(arr[i]))
		}
	}
	re := strings.Join(temp, " ")
	return re
}

func Test_reverseWords(t *testing.T) {
	reverseWords("the sky is blue")
}
