package 栈

import (
	"fmt"
	"testing"
)

/*
给定一个二叉树，返回它的中序 遍历。

https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/887/


中序遍历（LDR）：左，中，右
前序遍历（DLR）：中，左，右
后序遍历（LRD）：左，右，中

总结：遍历是以中为基准的，中在前则为前序遍历，中在中间则为中序遍历
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归遍历
func inorderTraversal(root *TreeNode) []int {
	val := make([]int, 0)
	var DFSTraversal func(node *TreeNode)
	DFSTraversal = func(node *TreeNode) {
		if node.Left != nil {
			DFSTraversal(node.Left)
		}
		val = append(val, node.Val)
		if node.Right != nil {
			DFSTraversal(node.Right)
		}
	}

	DFSTraversal(root)
	return val
}

// 堆栈遍历
func inorderTraversal2(root *TreeNode) []int {
	val := make([]int, 0)
	if root == nil {
		return val
	}

	visited := map[*TreeNode]bool{}

	stack := make([]*TreeNode, 0)

	popStack := func() *TreeNode {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}

	peekStack := func() *TreeNode {
		v := stack[len(stack)-1]
		return v
	}

	pushStack := func(v *TreeNode) {
		stack = append(stack, v)
	}

	pushStack(root)
	for len(stack) > 0 {
		node := peekStack()
		visited[node] = true
		if node.Left != nil {
			_, ok := visited[node.Left]
			if !ok {
				pushStack(node.Left)
				continue
			} else {
				_, ok := visited[node.Right]
				if ok {
					popStack()
					continue
				}
				val = append(val, node.Val)
			}

		} else {
			if node.Right != nil {
				_, ok := visited[node.Right]
				if ok {
					popStack()
					continue
				}
			}
			val = append(val, node.Val)
		}

		if node.Right != nil {
			_, ok := visited[node.Right]
			if !ok {
				pushStack(node.Right)
			}
			continue
		}
		popStack()
	}

	return val
}

// 网上代码
func inorderTraversal3(root *TreeNode) []int {
	val := make([]int, 0)
	if root == nil {
		return val
	}

	stack := make([]*TreeNode, 0)

	popStack := func() *TreeNode {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}

	pushStack := func(v *TreeNode) {
		stack = append(stack, v)
	}

	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			pushStack(cur)
			cur = cur.Left
		} else {
			cur = popStack()
			val = append(val, cur.Val)
			cur = cur.Right
		}
	}
	return val
}

func Test_inorderTraversal(t *testing.T) {
	root := &TreeNode{}
	node1 := &TreeNode{}
	node2 := &TreeNode{}

	root.Val = 1
	root.Left = nil
	root.Right = node1

	node1.Val = 2
	node1.Left = node2
	node1.Right = nil

	node2.Val = 3

	fmt.Println(inorderTraversal3(root))
}
