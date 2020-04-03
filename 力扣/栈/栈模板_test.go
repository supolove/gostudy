package 栈

import (
	"testing"
)

type MobanNode struct {
	Neighbor []*MobanNode
}

// 使用递归
// https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/882/
func moban1(cur *MobanNode, target *MobanNode, visited map[*MobanNode]bool) bool {
	if cur == target {
		return true
	}

	for _, next := range cur.Neighbor {
		_, ok := visited[next]
		if !ok {
			visited[next] = true
		}

		if moban1(next, target, visited) == true {
			return true
		}

	}
	return false
}

// 使用栈
// https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/886/
func moban2(root, target *MobanNode) bool {
	visited := map[*MobanNode]bool{}
	stack := make([]*MobanNode, 0)

	popStack := func() *MobanNode {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}

	pushStack := func(v *MobanNode) {
		stack = append(stack, v)
	}

	pushStack(root)
	for len(stack) > 0 {
		cur := popStack()
		if cur == target {
			return true
		}
		for _, next := range cur.Neighbor {
			_, ok := visited[next]
			if !ok {
				pushStack(next)
				visited[next] = true
			}
		}
	}

	return false
}

func Test_moban(t *testing.T) {

}
