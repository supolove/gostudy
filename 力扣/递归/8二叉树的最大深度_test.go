package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/259/complexity-analysis/1225/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memory []int

func digui5(root *TreeNode, n int) []int {
	if root.Left == nil && root.Right == nil {
		memory = append(memory, n)
		return memory
	}

	if root.Left != nil {
		digui5(root.Left, n+1)
	}

	if root.Right != nil {
		digui5(root.Right, n+1)
	}

	return memory
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	memory = []int{}
	m := digui5(root, 1)
	v := 0
	for i := 0; i < len(m); i++ {
		if v < m[i] {
			v = m[i]
		}
	}
	return v
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	if left < right {
		left = right
	}
	return left + 1
}

func Test_maxDepth(t *testing.T) {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 20}
	n4 := &TreeNode{Val: 15}
	n5 := &TreeNode{Val: 7}
	n1.Left = n2
	n1.Right = n3

	n3.Left = n4
	n3.Right = n5

	fmt.Println(maxDepth2(n1))
}
