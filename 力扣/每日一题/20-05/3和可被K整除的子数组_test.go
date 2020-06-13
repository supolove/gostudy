package _0_05

import (
	"fmt"
	"testing"
)

// 穷举法，超时
func subarraysDivByK(A []int, K int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			a := A[i : j+1]
			s := 0
			for m := 0; m < len(a); m++ {
				s += a[m]
			}
			if s%K == 0 {
				fmt.Println(a)
				res++
			}
		}
	}
	return res
}

/*
前缀和解决算法
P[i] = A[0] + A[1] + ... + A[i]
sum(i,j) = P[i] - P[i-1]
(P[i] - P[i-1]) mod K == 0
P[j] mod K == P[i-1] mod K

假设A【：i】 的和为 a，A【：j】的和为 b，并且 a % k == b % k，那么 A【j+1：i】的和必然能够整除 K。
那么你只需要记载在 i 之前，每个前缀和对 K 取余的数量即可。
*/
func subarraysDivByK2(A []int, K int) int {
	var m [10000]int
	m[0]++
	prefix, res := 0, 0
	for _, a := range A {
		prefix = (a + prefix) % K
		if prefix < 0 {
			prefix += K
		}
		res += m[prefix]
		m[prefix]++
	}
	return res
}

func Test_subarraysDivByK(t *testing.T) {
	fmt.Println(subarraysDivByK2([]int{4, 5, 0, -2, -3, 1}, 5))
}
