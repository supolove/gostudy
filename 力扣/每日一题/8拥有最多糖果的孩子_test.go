package 每日一题

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := 0
	for _, n := range candies {
		if n > max {
			max = n
		}
	}

	for i, n := range candies {
		if n+extraCandies >= max {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}

func Test_kidsWithCandies(t *testing.T) {
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
}
