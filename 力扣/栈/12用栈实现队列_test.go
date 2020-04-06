package 栈

import "testing"

/**
go没有堆栈和队列类库，所以只能这样实现，简单方便，但是不符合该题的题意
https://leetcode-cn.com/explore/learn/card/queue-stack/220/conclusion/888/
*/

type MyQueue struct {
	Stack []int
}

/** Initialize your data structure here. */
func Constructor1() MyQueue {
	queue := MyQueue{}
	queue.Stack = []int{}
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	v := this.Stack[0]
	this.Stack = this.Stack[1:]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Stack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack) == 0
}

func Test_queue(t *testing.T) {

}
