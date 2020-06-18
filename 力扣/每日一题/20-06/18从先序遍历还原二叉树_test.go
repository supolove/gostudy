package _0_06

import (
	"strconv"
	"testing"
)

// 使用栈实现
func recoverFromPreorder(S string) *TreeNode {
	var stack []string
	var nodes []*TreeNode

	push := func(value string) {
		stack = append(stack, value)
	}

	pop := func() string {
		value := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return value
	}

	peek := func() string {
		return stack[len(stack)-1]
	}

	s, _ := strconv.Atoi(string(S[0]))
	root := &TreeNode{Val: s}
	nodes = append(nodes, root)
	add := true
	for i := 1; i < len(S); i++ {
		if S[i] == '-' {
			for j := i; j < len(S); j++ {
				if S[j] != '-' {
					tag := S[i:j]
					if len(stack) > 0 && tag == peek() {
						add = false
						pop()
					} else {
						add = true
						push(S[i:j])
					}
					i = j - 1
					break
				}
			}
		} else {
			for j := i; j < len(S); j++ {
				if S[j] == '-' {
					if add {
						s, _ := strconv.Atoi(string(S[i:j]))
						n := &TreeNode{Val: s}
						nodes[len(nodes)-1].Left = n
						nodes = append(nodes, n)
					} else {
						nodes = nodes[0 : len(nodes)-1]
						s, _ := strconv.Atoi(string(S[i:j]))
						n := &TreeNode{Val: s}
						nodes[len(nodes)-1].Right = n
					}
					i = j - 1
					break
				}
			}
		}
	}
	return root
}

func recoverFromPreorder2(S string) *TreeNode {
	ans := map[int]*TreeNode{-1: &TreeNode{}}
	addTree := func(v, p int) {
		ans[p] = &TreeNode{Val: v}
		if ans[p-1].Left == nil {
			ans[p-1].Left = ans[p]
		} else {
			ans[p-1].Right = ans[p]
		}
	}
	val, dep, hasVal := 0, 0, false
	for _, c := range S {
		if c != '-' {
			val = 10*val + int(c-'0')
			hasVal = true
		} else if hasVal {
			addTree(val, dep)
			val, dep, hasVal = 0, 1, false
		} else {
			dep++
		}
	}
	addTree(val, dep)
	return ans[0]
}

func Test_recoverFromPreorder(t *testing.T) {
	recoverFromPreorder2("1-2--3--4-5--6--7")
}
