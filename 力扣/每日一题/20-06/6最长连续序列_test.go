package _0_06

import (
	"fmt"
	"testing"
)

func qSort(nums []int) {
	partition := func(arr []int, low, height int) int {
		base := arr[low]
		i := low
		j := height
		for i < j {
			for i < j && arr[j] >= base {
				j--
			}
			for i < j && arr[i] <= base {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[low], arr[i] = arr[i], arr[low]
		return i
	}

	var quickSort func(arr []int, low, height int)
	quickSort = func(arr []int, low, height int) {
		if low < height {
			p := partition(arr, low, height)
			quickSort(arr, low, p-1)
			quickSort(arr, p+1, height)
		}
	}

	quickSort(nums, 0, len(nums)-1)
}

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	qSort(nums)
	maxlen := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			if nums[i]+1 == nums[i+1] {
				index++
			} else if nums[i] == nums[i+1] {
				// 不处理
			} else {
				if maxlen < index {
					maxlen = index
				}
				index = 0
			}
		} else {
			if maxlen < index {
				maxlen = index
			}
			index = 0
		}
	}
	return maxlen + 1
}

func longestConsecutive2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	hash_dict := map[int]int{}
	max_length := 0
	for _, num := range nums {
		_, ok := hash_dict[num]
		if !ok {
			left := hash_dict[num-1]
			right := hash_dict[num+1]
			cur_length := 1 + left + right
			if cur_length > max_length {
				max_length = cur_length
			}
			hash_dict[num] = cur_length
			hash_dict[num-left] = cur_length
			hash_dict[num+right] = cur_length
		}
	}

	return max_length
}

func Test_longestConsecutive(t *testing.T) {
	num := []int{1, 2, 6, 0}
	//qSort(num)
	//fmt.Println(num)

	fmt.Println(longestConsecutive2(num))
}
