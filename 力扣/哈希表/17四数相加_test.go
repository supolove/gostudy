package 哈希表

import (
	"fmt"
	"testing"
)

// 超时
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	for a := 0; a < len(A); a++ {
		for b := 0; b < len(B); b++ {
			for c := 0; c < len(C); c++ {
				for d := 0; d < len(D); d++ {
					if A[a]+B[b]+C[c]+D[d] == 0 {
						res++
					}
				}
			}
		}
	}
	return res
}

func fourSumCount2(A []int, B []int, C []int, D []int) int {
	m := map[int]int{}
	res := 0
	for a := 0; a < len(A); a++ {
		for b := 0; b < len(B); b++ {
			m[A[a]+B[b]] = m[A[a]+B[b]] + 1
		}
	}

	for c := 0; c < len(C); c++ {
		for d := 0; d < len(D); d++ {
			n, ok := m[-C[c]-D[d]]
			if ok {
				res += n
			}
		}
	}
	return res
}

func Test_fourSumCount(t *testing.T) {
	fmt.Println(fourSumCount2([]int{2, 1}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
