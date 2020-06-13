package _0_05

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/symmetric-tree/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func digui(leftRoot *TreeNode, rightRoot *TreeNode) bool {
	if leftRoot == nil && rightRoot == nil {
		return true
	}
	if leftRoot == nil || rightRoot == nil {
		return false
	}
	if leftRoot.Val != rightRoot.Val {
		return false
	}
	return digui(leftRoot.Left, rightRoot.Right) && digui(leftRoot.Right, rightRoot.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return digui(root.Left, root.Right)
}

func Test_isSymmetric(t *testing.T) {
	//n1 := &TreeNode{Val:1}
	//n2 := &TreeNode{Val:2}
	//n3 := &TreeNode{Val:2}
	//n4 := &TreeNode{Val:3}
	//n5 := &TreeNode{Val:3}
	//n1.Left = n2
	//n1.Right = n3
	//
	//n2.Right = n4
	//n3.Right = n5

	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 2}
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3

	n2.Right = n4
	n3.Right = n5

	fmt.Println(isSymmetric(n1))

}
