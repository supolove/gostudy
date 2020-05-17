package 递归

import (
	"fmt"
	"testing"
)

func g(start, end int) []*TreeNode {
	var all []*TreeNode
	if start > end {
		all = append(all, nil)
		return all
	}

	for i := start; i <= end; i++ {
		left := g(start, i-1)
		right := g(i+1, end)
		for _, l := range left {
			for _, r := range right {
				current := &TreeNode{Val: i}
				current.Left = l
				current.Right = r
				all = append(all, current)
			}
		}
	}
	return all
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return g(1, n)
}

func Test_generateTrees(t *testing.T) {
	fmt.Println(generateTrees(5))
}
