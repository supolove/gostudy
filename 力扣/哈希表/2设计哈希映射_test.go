package 哈希表

import (
	"container/list"
	"testing"
)

type MyHashMap struct {
	Val [100]list.List
	Set [1000000]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		Set: [1000000]int{},
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if value == 0 {
		this.Set[key] = -1
		return
	}
	this.Set[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if this.Set[key] == -1 {
		return 0
	}

	if this.Set[key] == 0 {
		return -1
	}
	return this.Set[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.Set[key] = 0
}

func Test_(t *testing.T) {

}
