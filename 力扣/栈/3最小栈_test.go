package 栈

import (
	"fmt"
	"testing"
)

type MinStack struct {
	Data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Data: []int{}}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
}

func (this *MinStack) Pop() {
	this.Data = this.Data[:len(this.Data)-1]
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	min := this.Data[0]
	for _, i := range this.Data {
		if i < min {
			min = i
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func Test_MinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(5)
	obj.Push(2)
	obj.Push(9)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
