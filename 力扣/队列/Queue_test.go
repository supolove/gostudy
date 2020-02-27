package 队列

import (
	"fmt"
	"testing"
)

type MyCircularQueue struct {
	CircularQueue []int
	Head          int
	Tail          int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{
		CircularQueue: make([]int, k),
		Head:          0,
		Tail:          0,
	}
	for idx := range queue.CircularQueue {
		queue.CircularQueue[idx] = -1
	}
	return queue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {

	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.CircularQueue[this.Head] = value
		return true
	}

	// 余数除法，指针向后移一位
	this.Tail = (this.Tail + 1) % len(this.CircularQueue)
	this.CircularQueue[this.Tail] = value

	/*
		if this.Tail < len(this.CircularQueue) - 1 {
			this.Tail = this.Tail + 1
			this.CircularQueue[this.Tail] = value

		}else {
			this.Tail = 0
			this.CircularQueue[this.Tail] = value
		}*/

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.Head == this.Tail {
		this.CircularQueue[this.Head] = -1
		return true
	}

	// 余数除法，指针向后移一位
	this.CircularQueue[this.Head] = -1
	this.Head = (this.Head + 1) % len(this.CircularQueue)
	/*
		this.CircularQueue[this.Head] = -1
		this.Head = this.Head + 1
		if this.Head > len(this.CircularQueue) - 1 {
			this.Head = 0
		}*/
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	return this.CircularQueue[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	return this.CircularQueue[this.Tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.CircularQueue[this.Head] == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.Head == 0 {
		return this.Tail == len(this.CircularQueue)-1
	}
	return this.Tail+1 == this.Head
}

func TestQuue(t *testing.T) {
	queue := Constructor(6)
	fmt.Println(queue.EnQueue(6))
	fmt.Println(queue.Rear())
	fmt.Println(queue.Rear())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(5)) //wrong
	fmt.Println(queue.Rear())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Front())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	//
	//
	//
	//
	//queue.EnQueue(10)
	//queue.EnQueue(12)
	//queue.EnQueue(13)
	//
	//fmt.Println(queue.DeQueue())
	//fmt.Println(queue.EnQueue(14))
	//
	//fmt.Println("IsEmpty:" ,queue.IsEmpty())
	//fmt.Println("IsFull:" ,queue.IsFull())
	//fmt.Println("Front:" , queue.Front())
	//fmt.Println("Tail:" , queue.Rear())

}
