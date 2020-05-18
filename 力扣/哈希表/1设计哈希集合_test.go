package 哈希表

import (
	"fmt"
	"testing"
)

/*
type MyHashSet struct {
	M [1000000]int
}


func Constructor() MyHashSet {
	return MyHashSet{M: [1000000]int{}}
}


func (this *MyHashSet) Add(key int)  {
	//if !this.Contains(key) {
	//	this.M = append(this.M, key)
	//}
	if key == 0 {
		this.M[key] = -1
	}else {
		this.M[key] = key
	}

}


func (this *MyHashSet) Remove(key int)  {
	//for i:=0; i< len(this.M) ;i++  {
	//	if this.M[i] == key {
	//		this.M = append( this.M[:i], this.M[i+1:]...)
	//	}
	//}
	this.M[key] = 0
}


func (this *MyHashSet) Contains(key int) bool {
	//for i:=0; i< len(this.M) ;i++  {
	//	if this.M[i] == key {
	//		return true
	//	}
	//}
	//return false
	return this.M[key] != 0 || this.M[key] == -1
}
*/

type MyHashSet struct {
	set []byte
}

/** Initialize your data structure here. */
func Constructor1() MyHashSet {
	return MyHashSet{
		set: make([]byte, 100),
	}
}

func (this *MyHashSet) Add(key int) {
	if key/8 >= len(this.set) {
		expandLen := ((key / 8) - len(this.set)) + 1
		if expandLen > 1000000/8+1 {
			expandLen = 1000000/8 + 1
		}

		this.set = append(this.set, make([]byte, expandLen)...)
	}

	this.set[key/8] = (1 << uint(key%8)) | this.set[key/8]

	// fmt.Println("add", key, this.Contains(key))
}

func (this *MyHashSet) Remove(key int) {
	if key/8 < len(this.set) {
		this.set[key/8] = ^(1 << uint(key%8)) & this.set[key/8]
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	exist := false
	if key/8 >= len(this.set) {
		exist = false
	} else if (1<<uint(key%8))&this.set[key/8] == 1<<uint(key%8) {
		exist = true
	}
	return exist
}

func Test_MyHashSet(t *testing.T) {
	obj := Constructor1()
	obj.Add(1)
	obj.Add(2)
	fmt.Println(obj.Contains(1))
	fmt.Println(obj.Contains(3))
	obj.Add(2)
	obj.Remove(2)
	fmt.Println(obj.Contains(2))
}
