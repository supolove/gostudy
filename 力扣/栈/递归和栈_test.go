package 栈

import (
	"fmt"
	"testing"
)

/*
	改项不属于力扣的题目，只是自己为了再深入学习一下递归和栈
*/

// 如果没有终结条件，将会内存爆满，程序将会挂掉
// fatal error: stack overflow

func recursive(n int) int {
	fmt.Println(n)

	// 添加终结条件，使递归终止
	if n == 5 {
		return n
	}

	// 先计算递归参数，后计算return操作
	return recursive(n+1) + 1
}

// 兔子问题
// 1 1 2 3 5 8 13 ... n
// 计算第n个的数

func tuzi(n int, m int, idx int) int {
	fmt.Println(m)

	// 添加终结条件，使递归终止
	if m == 13 {
		return idx
	}

	return tuzi(m, n+m, idx+1)
}

// for 循环兔子问题
func tuzi2() int {
	//f(n) = f(n) + f(n+1)
	i := 0
	n := 1
	m := 1
	for i < 10 {
		i++
		t := m
		m = n + m
		n = t
		fmt.Println(m)
	}
	return m
}

func Test_Recursive(t *testing.T) {
	//fmt.Println(recursive(2))
	//fmt.Println(tuzi(1,1, 1))
	fmt.Println(tuzi2())
}
