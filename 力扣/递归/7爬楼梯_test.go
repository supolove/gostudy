package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/258/memorization/1213/


解题思路：

1. 递归：走到n层台阶的的方法取决于到n-1和n-2层的方法之和，而到n-1层的方法又取决于到n-2层和n-3层的方法之和，
  直至到第2层和第1层，到第2层有两种方法，到第1层有1种方法，所以递归求和即可

2. 动态规划：由于递归过程会涉及到大量的重复计算，因为到具体某一层的方法数目实际是确定的，例如到1层只有1种方法，
到2层只有两种方法，到3层有3种方法，到4层的方法数由到2层和3层的方法数决定，到n层的方法数由到n-1层和n-2层确定

阶   几种方法
1 -> 1
2 -> 2
3 -> 3
4 -> 5
5 -> 8

*/

//var cache2 map[int]int
var cache2 = map[int]int{}

func climbStairs(n int) int {
	v, ok := cache2[n]
	if ok {
		return v
	}
	if n == 0 || n == 1 {
		cache2[0] = 1
		cache2[1] = 1
		return 1
	}
	sumPath1 := climbStairs(n - 1)
	sumPath2 := climbStairs(n - 2)
	result := sumPath1 + sumPath2
	cache2[n] = result
	return result
}

// 没使用缓存会重复计算
func climbStairs2(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	sumPath1 := climbStairs(n - 1)
	sumPath2 := climbStairs(n - 2)
	result := sumPath1 + sumPath2
	return result
}

func Test_climbStairs(t *testing.T) {
	fmt.Println(climbStairs(5))
}
