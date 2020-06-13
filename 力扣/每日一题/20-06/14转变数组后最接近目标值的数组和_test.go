package _0_06

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/
*/

func findBestValue(arr []int, target int) int {
	// 先排序
	sort.Ints(arr)

	l := len(arr)
	// 参考值
	ref := target / l

	// 在数组中找到参考值
	left := 0
	right := l - 1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if ref == arr[mid] {
			break
		} else if ref > arr[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == right {
		mid = left
	}

	sum := 0
	for i := 0; i < mid; i++ {
		sum += arr[i]
	}

	value := arr[mid]
	distence := 100000
	for {
		s := sum + (l-mid)*value
		if int(math.Abs(float64(target-s))) == 0 {
			return value
		} else if int(math.Abs(float64(target-s))) > distence {
			return value - 1
		} else {
			distence = int(math.Abs(float64(target - s)))
			value++
		}
	}
}

func findBestValue2(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if target > arr[0] {
			return arr[0]
		}
		return target
	}
	dest := target / n
	fmt.Println(target % n)
	if target%n > n/2 {
		dest += 1
	}
	fmt.Println(dest)
	l := make([]int, 0)
	r := make([]int, 0)
	for i := 0; i < n; i++ {
		if arr[i] <= dest {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}
	if len(l) == 0 {
		return dest
	}
	return findBestValue2(r, target-sum(l))
}

func sum(l []int) int {
	ret := 0
	for i := 0; i < len(l); i++ {
		ret += l[i]
	}
	return ret
}

func Test_findBestValue(t *testing.T) {
	fmt.Println(findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803))
}
