package _0_06

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
*/

var ret = math.MinInt32

func maxPathSum(root *TreeNode) int {
	/**
	  对于任意一个节点, 如果最大和路径包含该节点, 那么只可能是两种情况:
	  1. 其左右子树中所构成的和路径值较大的那个加上该节点的值后向父节点回溯构成最大路径
	  2. 左右子树都在最大路径中, 加上该节点的值构成了最终的最大路径
	  **/
	ret = math.MinInt32
	getMax(root)
	return ret
}

func getMax(r *TreeNode) int {
	if r == nil {
		return 0
	}
	left := maxInt(0, getMax(r.Left)) // 如果子树路径和为负则应当置0表示最大路径不包含子树
	right := maxInt(0, getMax(r.Right))
	ret = maxInt(ret, r.Val+left+right) // 判断在该节点包含左右子树的路径和是否大于当前最大路径和
	return maxInt(left, right) + r.Val
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maxPathSum(t *testing.T) {
	n1 := &TreeNode{Val: -10}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 20}
	n4 := &TreeNode{Val: 15}
	n5 := &TreeNode{Val: 7}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5

	fmt.Println(maxPathSum(n1))
}
