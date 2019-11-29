package 数据结构

import (
	"container/ring"
	"fmt"
	"testing"
)

// 环

func TestRing(t *testing.T) {
	r := ring.New(3)

	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}

	// 计算 1+2+3
	s := 0
	r.Do(func(p interface{}) {
		s += p.(int)
	})

	r.Do(func(p interface{}) {
		s += p.(int)
	})

	fmt.Println("sum is", s)
}
