package 二分查找

import (
	"fmt"
	"math"
	"testing"
)

func mySqrt(x int) int {
	a := math.Sqrt(float64(x))
	return int(a)
}

func mySqrt2(x int) int {
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		pow := mid * mid
		if pow == x {
			return mid
		}
		if pow > x {
			right = mid - 1
		} else {
			left = mid + 1
			if left*left > x {
				return left - 1
			}
		}
	}
	return 0
}

func Test_mySqrt(t *testing.T) {
	fmt.Println(mySqrt2(9))
}
