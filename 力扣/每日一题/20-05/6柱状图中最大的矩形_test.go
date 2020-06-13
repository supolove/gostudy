package _0_05

import (
	"fmt"
	"testing"
)

/**
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
困难

相邻的柱子

看了别人的答案想了半天才明白……其实可以把这个想象成锯木板，如果木板都是递增的那我很开心，
如果突然遇到一块木板i矮了一截，那我就先找之前最戳出来的一块（其实就是第i-1块），
计算一下这个木板单独的面积，然后把它锯成次高的，这是因为我之后的计算都再也用不着这块木板本身的高度了。
再然后如果发觉次高的仍然比现在这个i木板高，那我继续单独计算这个次高木板的面积（应该是第i-1和i-2块），
再把它俩锯短。直到发觉不需要锯就比第i块矮了，那我继续开开心心往右找更高的。当然为了避免到了最后一直都是递增的，
所以可以在最后加一块高度为0的木板。

这个算法的关键点是把那些戳出来的木板早点单独拎出来计算，然后就用不着这个值了。说实话真的很佩服第一个想出来的人……

*/

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

/*
使用 left 和 right 数组记录以 i 柱子为基准，从 i 开始向左右两边延伸的第一根比 i 小的柱子
那么，在 （left[i], right[i]) (注意是小括号，即不包括边界)之间的柱子都是比 i 柱子小的
那么 i 柱子能够围成的面积宽度是 width = right[i] - left[i] - 1
面积则为 heights[i] * width
*/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	left := make([]int, n)
	right := make([]int, n)

	//第一根柱子，左边不存在比它小的
	left[0] = -1
	//最后一根柱子，右边不存在比它小的
	right[n-1] = n

	for i := 0; i < n; i++ {
		temp := i - 1
		for temp >= 0 && heights[temp] >= heights[i] {
			temp = left[temp]
		}
		//当上述循环 break 后，  temp 即为左边第一根小于 i 位置的柱子
		left[i] = temp
	}
	for i := n - 2; i >= 0; i-- {
		temp := i + 1
		for temp < n && heights[temp] >= heights[i] {
			temp = right[temp]
		}
		//当上述循环 break 后，  temp 即为左边第一根小于 i 位置的柱子
		right[i] = temp
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1
		maxArea = max(maxArea, heights[i]*width)
	}
	return maxArea
}

func largestRectangleArea2(heights []int) int {
	ans, s, hs := 0, []int{0}, make([]int, len(heights)+2)
	copy(hs[1:len(heights)+1], heights)
	for i, h := range hs {
		for n := len(s); hs[s[n-1]] > h; n = len(s) {
			if area := (i - 1 - s[n-2]) * hs[s[n-1]]; area > ans {
				ans = area
			}
			s = s[:n-1]
		}
		s = append(s, i)
	}
	return ans
}

// 单调栈
func largestRectangleArea3(heights []int) int {
	return 0
}

func Test_largestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
