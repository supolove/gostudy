package 栈

import (
	"fmt"
	"testing"
)

/**
go没有堆栈和队列类库，所以只能这样实现，简单方便，但是不符合该题的题意
https://leetcode-cn.com/explore/learn/card/queue-stack/220/conclusion/889/
*/

type MyStack struct {
	Queue []int
}

/** Initialize your data structure here. */
func Constructor3() MyStack {
	stack := MyStack{}
	stack.Queue = []int{}
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Queue[len(this.Queue)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}

func Test_stack(t *testing.T) {
	s := Constructor3()
	s.Push(1)
	s.Pop()
	fmt.Println(s.Empty())
}
