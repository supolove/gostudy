package 哈希表

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/207/conclusion/831/
*/

type RandomizedSet struct {
	M map[int]bool
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{M: map[int]bool{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.M[val]; ok {
		return false
	}
	this.M[val] = true
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.M[val]; ok {
		delete(this.M, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(this.M))
	i := 0
	for k := range this.M {
		if i == r {
			return k
		}
		i++
	}
	return 2
}

func Test_RandomizedSet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	fmt.Println(i)
}
