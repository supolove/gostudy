package 队列

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

/*

说明：
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。


示例1：
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。


示例2：
输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。


示例 3:
输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。


示例 4:
输入: deadends = ["0000"], target = "8888"
输出：-1


提示：
死亡列表 deadends 的长度范围为 [1, 500]。
目标数字 target 不会在 deadends 之中。
每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
*/

// 旋转
func turn() {

}

//

func openLock(deadends []string, target string) int {
	// 优化把[]string数组转换为4个Int类型数组
	arr1 := []int{}
	arr2 := []int{}
	arr3 := []int{}
	arr4 := []int{}
	/*
		appendArr := func(arr []int, vs string) {
			v,_:= strconv.Atoi(vs)
			i := sort.SearchInts(arr, v)
			if !(i < len(arr) && arr[i] == v) {
				arr = append(arr, v)
				sort.IntsAreSorted(arr)
			}
		}*/

	for _, ss := range deadends {
		//appendArr(arr1, ss[0:1])
		//appendArr(arr2, ss[1:2])
		//appendArr(arr3, ss[2:3])
		//appendArr(arr4, ss[3:4])

		v1, _ := strconv.Atoi(ss[0:1])
		sort1 := sort.SearchInts(arr1, v1)
		if !(sort1 < len(arr1) && arr1[sort1] == v1) {
			arr1 = append(arr1, v1)
			sort.Ints(arr1)
		}

		v2, _ := strconv.Atoi(ss[1:2])
		sort2 := sort.SearchInts(arr2, v2)
		if !(sort2 < len(arr2) && arr2[sort2] == v2) {
			arr2 = append(arr2, v2)
			sort.Ints(arr2)
		}

		v3, _ := strconv.Atoi(ss[2:3])
		sort3 := sort.SearchInts(arr3, v3)
		if !(sort3 < len(arr3) && arr3[sort3] == v3) {
			arr3 = append(arr3, v3)
			sort.Ints(arr3)
		}

		v4, _ := strconv.Atoi(ss[3:4])
		sort4 := sort.SearchInts(arr4, v4)
		if !(sort4 < len(arr4) && arr4[sort4] == v4) {
			arr4 = append(arr4, v4)
			sort.Ints(arr4)
		}
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	return 0
}

func TestZhuanpan(t *testing.T) {
	arr := []string{"0201", "0101", "1212", "2002"}
	target := "0202"
	openLock(arr, target)
}
