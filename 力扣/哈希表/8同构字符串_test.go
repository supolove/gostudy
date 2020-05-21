package 哈希表

import (
	"fmt"
	"reflect"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/205/practical-application-hash-map/812/

题目有点难理解
意思是比如：egg，格式就相当于 122，add也是122,所以返回true
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	var sli []int
	i := 0
	for _, b := range s {
		v, ok := m[b]
		if ok {
			sli = append(sli, v)
		} else {
			i++
			m[b] = i
			sli = append(sli, i)
		}
	}

	m = map[rune]int{}
	var sli2 []int
	i = 0
	for _, b := range t {
		v, ok := m[b]
		if ok {
			sli2 = append(sli2, v)
		} else {
			i++
			m[b] = i
			sli2 = append(sli2, i)
		}
	}
	return reflect.DeepEqual(sli, sli2)
}

func isIsomorphic2(s string, t string) bool {

	return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}

func isIsomorphicHelper(s string, t string) bool {
	length := len(s)
	cache := make(map[byte]byte)

	for i := 0; i < length; i++ {
		if _, exists := cache[s[i]]; exists {
			if cache[s[i]] != t[i] {
				return false
			}
		} else {
			cache[s[i]] = t[i]
		}
	}

	return true
}

func Test_isIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("aba", "baa"))
}
