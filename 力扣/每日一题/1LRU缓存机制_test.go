package 每日一题

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/lru-cache/
2020-05-25
*/

type LRUCache struct {
	arr []int
	m   map[int]int
	c   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{arr: []int{}, m: map[int]int{}, c: capacity}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if ok {
		for i := 0; i < len(this.arr); i++ {
			if this.arr[i] == key {
				if i == 0 {
					this.arr = append(this.arr[1:], key)
				} else if i == len(this.arr)-1 {
					this.arr = append(this.arr[:i], key)
				} else {
					this.arr = append(this.arr[:i], this.arr[i+1:]...)
					this.arr = append(this.arr, key)
				}
				break
			}
		}
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.m[key]
	if ok {
		this.m[key] = value
		for i := 0; i < len(this.arr); i++ {
			if this.arr[i] == key {
				if i == 0 {
					this.arr = append(this.arr[1:], key)
				} else if i == len(this.arr)-1 {
					this.arr = append(this.arr[:i], key)
				} else {
					this.arr = append(this.arr[:i], this.arr[i+1:]...)
					this.arr = append(this.arr, key)
				}
				break
			}
		}
		return
	}
	if len(this.arr) > this.c-1 {
		delete(this.m, this.arr[0])
		this.arr = this.arr[1:]
	}
	this.arr = append(this.arr, key)
	this.m[key] = value
}

func Test_lru(t *testing.T) {
	o := Constructor(3)
	o.Put(1, 1)
	o.Put(2, 2)
	o.Put(3, 3)
	o.Put(4, 4)
	fmt.Println(o.Get(4))
	fmt.Println(o.Get(3))
	fmt.Println(o.Get(2))
	fmt.Println(o.Get(1))
	o.Put(5, 5)
	fmt.Println(o.Get(1))
	fmt.Println(o.Get(2))
	fmt.Println(o.Get(3))
	fmt.Println(o.Get(4))
	fmt.Println(o.Get(5))

	//fmt.Println(o.Get(1))
	//o.Put(3,3)
	//fmt.Println(o.Get(2))
	//o.Put(4,4)
	//fmt.Println(o.Get(1))
	//fmt.Println(o.Get(3))
	//fmt.Println(o.Get(4))
	//a := []int{1,2,3,4,5}
	//fmt.Println(a[:4])
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(append(a[:2], a[3:]...))
}
