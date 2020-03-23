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
	if target == "0000" {
		return -1
	}
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
		mv := v[0:idx] + strconv.Itoa(v1) + v[idx+1:]
		return mv
	}

	addValue := func(v string, idx int) string {
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

	// 层次，值
	type VP struct {
		Floor int
		Value string
	}

	start := "0000"
	m[start] = 0
	var q []VP
	q = append(q, VP{0, start})
	for {
		if len(q) > 0 {
			item := q[0]
			//fmt.Println(item)
			if item.Value == target {
				return item.Floor
			}
			q = q[1:]
			for i := 0; i < 4; i++ {
				r := addValue(item.Value, i)
				if !isInDeadends(r) {
					q = append(q, VP{item.Floor + 1, r})
					m[r] = 0
				}
			}
			for i := 0; i < 4; i++ {
				r := reduceValue(item.Value, i)
				if !isInDeadends(r) {
					q = append(q, VP{item.Floor + 1, r})
					m[r] = 0
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
	result := openLock2(arr, target)
	fmt.Println("result = ", result)
}

/**
网友最短时间解法
*/
//官方解法只有一种 bfs

func openLock2(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	visited := make(map[string]bool)
	for _, v := range deadends {
		if v == "0000" {
			return -1
		}
		visited[v] = true
	}

	// bfs --------------------------------------------------------------
	startQueue := make(map[string]bool) // 构造处理字符串队列。用于从起点侧开始搜索
	startQueue["0000"] = true           // 起点
	endQueue := make(map[string]bool)   // 构造处理字符串队列。用于从终点侧开始搜索
	endQueue[target] = true

	return bfs(startQueue, endQueue, visited, 0) // count从0开始
}

//81% 100%
// bfs. 用BFS模拟了双向搜索的步骤。 count为第几步
func bfs(start, end, visited map[string]bool, count int) int {
	if len(start) <= 0 {
		return -1
	} // 出现断层 (就是start这边搜索不到过去target的路径了，其队列就没有东西存着)
	if len(start) > len(end) { // 从小的那一端开始
		return bfs(end, start, visited, count)
	}

	change := []uint8{9, 1}        // 转动数字的增量。9代表向后反转。注意要是uint8（byte是它的别名）
	nexts := make(map[string]bool) //存储start端搜索的下一步需要处理的状态点字符串
	var curSlice []byte
	var origin byte
	var nextStr string // 下一步的状态（字符串）

	// 处理start队列（从队头到队尾，这由slice遍历机制决定）
	for cur := range start {
		visited[cur] = true      // 标记为已访问
		curSlice = []byte(cur)   // 字符串转为[]byte
		for i := 0; i < 4; i++ { // 遍历四位数字（四个轮盘状态）
			origin = curSlice[i]     // 备份当前字符
			for j := 0; j < 2; j++ { // 正反转动
				curSlice[i] = (curSlice[i]-'0'+change[j])%10 + '0'
				nextStr = string(curSlice)
				if !visited[nextStr] {
					if _, ok := end[nextStr]; ok { // end队列也有，说明碰撞了，下一步就可以见面
						count++
						return count
					} else {
						nexts[nextStr] = true
					}
				}
				curSlice[i] = origin // 复原单词
			}
		}
	}
	count++
	return bfs(nexts, end, visited, count)
}
