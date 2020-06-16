package _0_06

import (
	"strconv"
	"strings"
	"testing"
)

/*
https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	Root *TreeNode
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.Root = root
	var queue []*TreeNode
	queue = append(queue, root)
	res := ""
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root != nil {
			res += strconv.Itoa(root.Val)
			queue = append(queue, root.Left)
			queue = append(queue, root.Right)
		} else {
			res += "n"
		}
		res += " "
	}
	//fmt.Println(res)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tree := strings.Fields(data)
	print(tree)
	if tree[0] == "n" {
		return nil
	}
	var queue []*TreeNode
	v, _ := strconv.Atoi(tree[0])
	root := &TreeNode{Val: v}
	queue = append(queue, root)
	i := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			continue
		}

		if tree[i] != "n" {
			v, _ = strconv.Atoi(tree[i])
			cur.Left = &TreeNode{Val: v}
		} else {
			cur.Left = nil
		}
		if tree[i+1] != "n" {
			v, _ = strconv.Atoi(tree[i+1])
			cur.Right = &TreeNode{Val: v}
		} else {
			cur.Right = nil
		}
		i += 2
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return root
}

func Test_treenode(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n1.Left = n2
	n1.Right = n3

	n3.Left = n4
	n3.Right = n5

	c := Codec{}
	c.Root = n1
	res := c.serialize(n1)

	c.deserialize(res)
}
