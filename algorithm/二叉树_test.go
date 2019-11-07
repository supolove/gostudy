package algorithm

import (
	"fmt"
	"golang.org/x/tour/tree"
	"testing"
)

/*
比较二叉平衡树内容是否相同
这道题可以练习一下二叉树的遍历、channel的状态测试等知识点
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkInternal(t, ch)
	close(ch)
}

func WalkInternal(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkInternal(t.Left, ch)
		ch <- t.Value
		WalkInternal(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if ok1 == false {
			break
		}
	}

	return true
}

func TestTree(t *testing.T) {
	c := make(chan string)
	go func() {
		c <- nil
	}()

	a, b := <-c
	fmt.Println(a, b)

	//ch := make(chan int)
	//go Walk(tree.New(2), ch)
	//for i := range ch{
	//	fmt.Println(i)
	//}
	//fmt.Println(Same(tree.New(1), tree.New(2)))
	//fmt.Println(Same(tree.New(1), tree.New(2)))
}
