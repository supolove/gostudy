package 递归

import (
	"fmt"
	"math/bits"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/260/conclusion/1231/
*/

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	kk := 0
	if K%2 > 0 {
		kk = K/2 + 1
	} else {
		kk = K / 2
	}
	r := kthGrammar(N-1, kk)
	if (r == 1 && K%2 > 0) || (r == 0 && K%2 == 0) {
		return 1
	} else {
		return 0
	}
}

// 精要处理，异或相同为0，不同为1
func kthGrammar2(N int, K int) int {
	if N == 1 {
		return 0
	}
	return (1 - K%2) ^ kthGrammar2(N-1, (K+1)/2)
}

// 二进制计数
func kthGrammar3(N int, K int) int {
	return bits.OnesCount(uint(K-1)) % 2
}

func Test_kthGrammar(t *testing.T) {
	fmt.Println(kthGrammar3(4, 5))
}
