package 递归

import (
	"fmt"
	"testing"
)

func reverse(s []byte, n int) {
	if n > len(s)/2-1 {
		return
	}
	s[n], s[len(s)-n-1] = s[len(s)-n-1], s[n]
	reverse(s, n+1)
}
func reverseString(s []byte) {
	reverse(s, 0)
	fmt.Println(s)
}

func Test_reverseString(t *testing.T) {
	b := []byte{'h'}
	reverseString(b)
}
