package _0_05

import (
	"fmt"
	"math"
	"testing"
)

// 只比较了本num和下一个num，有问题
func rob(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			sum += nums[i]
			break
		}
		if i > len(nums)-1 {
			break
		}
		if nums[i] > nums[i+1] {
			sum += nums[i]
		} else {
			sum += nums[i+1]
			i++
		}
		i++
	}
	return sum
}

// 动态规划
func rob2(nums []int) int {
	last, now := 0, 0
	for _, n := range nums {
		max := int(math.Max(float64(last+n), float64(now)))
		last, now = now, max
	}
	return now
}

func Test_rob(t *testing.T) {

	// 2 1 1 2, 2 1 1 2 3 1
	//fmt.Println(rob([]int{2,1,2,3,5,1}))
	fmt.Println(rob2([]int{2, 1, 2, 3, 5, 1}))
}
