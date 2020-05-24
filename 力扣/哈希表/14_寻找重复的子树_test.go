package 哈希表

import (
	"strconv"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/206/practical-application-design-the-key/823/
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(m map[int]*TreeNode, node *TreeNode, res *[]*TreeNode) {

	if node.Left != nil {
		dfs(m, node.Left, res)
	}

	if node.Right != nil {
		dfs(m, node.Right, res)
	}

	n, ok := m[node.Val]
	if ok {
		cur := node
		mn := n
		for {
			//if mn.Val != cur.Val{
			//	break
			//}
			if mn.Left != nil && cur.Left != nil {
				mn = mn.Left
				cur = cur.Left
				continue
			}

			if mn.Right != nil && cur.Right != nil {
				mn = mn.Right
				cur = cur.Right
				continue
			}

			//　如果遍历道叶子，说明两个子树相等
			if mn.Left == nil && mn.Right == nil && cur.Left == nil && cur.Right == nil {
				*res = append(*res, n)
				delete(m, n.Val)
				return
			}
			return
		}
	}

	m[node.Val] = node
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := map[int]*TreeNode{}
	var res []*TreeNode
	dfs(m, root.Left, &res)
	dfs(m, root.Right, &res)
	return res
}

func dfs2(root *TreeNode, ans *[]*TreeNode, mp map[string]int) string {
	if root == nil {
		return "#"
	}

	s := strconv.Itoa(root.Val)
	tmp := s + "," + dfs2(root.Left, ans, mp) + dfs2(root.Right, ans, mp)
	if mp[tmp] == 1 {
		*ans = append(*ans, root)
	}
	mp[tmp] += 1
	return tmp
}

func findDuplicateSubtrees2(root *TreeNode) []*TreeNode {
	m := map[string]int{}
	var res []*TreeNode
	dfs2(root, &res, m)
	return res
}

func Test_findDuplicateSubtrees(t *testing.T) {
	//n1 := &TreeNode{Val:1}
	//n2 := &TreeNode{Val:2}
	//n3 := &TreeNode{Val:4}
	//n4 := &TreeNode{Val:3}
	//n5 := &TreeNode{Val:2}
	//n6 := &TreeNode{Val:4}
	//n7 := &TreeNode{Val:4}
	//
	//n1.Left = n2
	//n2.Left = n3
	//
	//n1.Right = n4
	//n4.Left = n5
	//n5.Left = n6
	//n4.Right = n7

	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 2}
	n5 := &TreeNode{Val: 3}
	n1.Left = n2
	n2.Left = n3
	n1.Right = n4
	n4.Left = n5

	//n1 := &TreeNode{Val:0}
	//n2 := &TreeNode{Val:0}
	//n3 := &TreeNode{Val:0}
	//n4 := &TreeNode{Val:0}
	//n5 := &TreeNode{Val:0}
	//n6 := &TreeNode{Val:0}
	//n1.Left = n2
	//n2.Left = n3
	//n1.Right = n4
	//n4.Right = n5
	//n5.Right = n6

	//n1 := &TreeNode{Val:2}
	//n2 := &TreeNode{Val:1}
	//n3 := &TreeNode{Val:11}
	//n4 := &TreeNode{Val:11}
	//n5 := &TreeNode{Val:1}
	//n1.Left = n2
	//n2.Left = n3
	//n1.Right = n4
	//n4.Left = n5

	findDuplicateSubtrees2(n1)

}
