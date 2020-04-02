package 栈

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/884/
*/

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Neighbors []*Node
* }
 */

//Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	if len(node.Neighbors) == 0 {
		return &Node{node.Val, []*Node{}}
	}

	m := map[*Node]*Node{}
	return dfsClone(node, m)
}

func dfsClone(node *Node, visited map[*Node]*Node) *Node {
	_, ok := visited[node]
	if ok {
		return visited[node]
	}

	p := &Node{node.Val, []*Node{}}
	visited[node] = p

	for _, n := range node.Neighbors {
		p.Neighbors = append(p.Neighbors, dfsClone(n, visited))
	}
	return p
}

func Test_cloneGraph(t *testing.T) {
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}
	n4 := &Node{}

	n1.Val = 1
	n2.Val = 2
	n3.Val = 3
	n4.Val = 4

	neighbor1 := []*Node{}
	neighbor2 := []*Node{}
	neighbor3 := []*Node{}
	neighbor4 := []*Node{}

	neighbor1 = append(neighbor1, n2, n4)
	neighbor2 = append(neighbor2, n1, n3)
	neighbor3 = append(neighbor3, n2, n4)
	neighbor4 = append(neighbor4, n3, n1)

	n1.Neighbors = neighbor1
	n2.Neighbors = neighbor2
	n3.Neighbors = neighbor3
	n4.Neighbors = neighbor4

	cloneGraph(n1)
}
