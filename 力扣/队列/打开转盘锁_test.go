package 队列

import (
	"fmt"
	"strconv"
	"testing"
)

/**
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


解题思路:
使用广度搜索算法
自己不会写，从解中copy了一个过来，基本思路是从0000开始找，先找一步能走到的所有
（1000，9000，0100,0900,0010,0090,0001,0009），加到队列里，
如果这中间找到了target，直接返回步数，
如果碰到了deadend中的一个，则退出循环继续找，相当于不把deadend中的入队，也就不会接着这一个继续找，
从而路径中不会包含带deadend中的坐标；那么如果找到的这一个不再deadend中，就把他添加到队列中，
在下一层的循环中接着找。直到找到target，这就是最短的路径。也就是BFS广度优先搜索。
*/

func openLock(deadends []string, target string) int {
	m := make(map[string]interface{})
	for _, i := range deadends {
		m[i] = 0
	}

	isInDeadends := func(item string) bool {
		_, ok := m[item]
		return ok
	}

	makeValue := func(v string, v1 int, idx int) string {
		if idx == 0 {
			return strconv.Itoa(v1) + v[idx+1:]
		}

		if idx == 3 {
			return v[:idx] + strconv.Itoa(v1)
		}

		return v[idx-1:idx] + strconv.Itoa(v1) + v[idx:]
	}

	addValue := func(v string, idx int) string {
		fmt.Println("str=", v, "idx=", idx)
		v1, _ := strconv.Atoi(v[idx : idx+1])
		if v1 != 9 {
			v1++
		} else {
			v1 = 0
		}
		return makeValue(v, v1, idx)
	}
	reduceValue := func(v string, idx int) string {
		v1, _ := strconv.Atoi(v[idx : idx+1])
		if v1 != 0 {
			v1--
		} else {
			v1 = 9
		}
		return makeValue(v, v1, idx)
	}

	start := "0000"
	var q []string
	q = append(q, start)
	nums := 0
	for {
		if len(q) > 0 {
			for idx := range q {
				item := q[idx]
				nums++
				if item == target {
					return nums
				}
				q = q[1:]
				for i := 0; i < 4; i++ {
					r := addValue(item, i)
					if !isInDeadends(r) {
						q = append(q, r)
					}
				}
				for i := 0; i < 4; i++ {
					r := reduceValue(item, i)
					if !isInDeadends(r) {
						q = append(q, r)
					}
				}
			}
		} else {
			return -1
		}
	}
}

func TestZhuanpan(t *testing.T) {
	arr := []string{
		"0201",
		"0101",
		"1212",
		"2002"}
	target := "0202"
	openLock(arr, target)
}
