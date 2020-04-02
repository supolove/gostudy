package 栈

import (
	"fmt"
	"strconv"
	"testing"
)

/*
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：
	整数除法只保留整数部分。
	给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

示例 1：
	输入: ["2", "1", "+", "3", "*"]
	输出: 9
	解释: ((2 + 1) * 3) = 9


示例 2：
	输入: ["4", "13", "5", "/", "+"]
	输出: 6
	解释: (4 + (13 / 5)) = 6

示例 3：
	输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
	输出: 22
	解释:
	  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
	= ((10 * (6 / (12 * -11))) + 17) + 5
	= ((10 * (6 / -132)) + 17) + 5
	= ((10 * 0) + 17) + 5
	= (0 + 17) + 5
	= 17 + 5
	= 22
*/

func evalRPN(tokens []string) int {
	stack := make([]string, 0)

	popStack := func() string {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}

	pushStack := func(v string) {
		stack = append(stack, v)
	}

	for _, v := range tokens {

		switch v {
		case "+":
			v1, _ := strconv.Atoi(popStack())
			v2, _ := strconv.Atoi(popStack())
			pushStack(strconv.Itoa(v1 + v2))
		case "-":
			v1, _ := strconv.Atoi(popStack())
			v2, _ := strconv.Atoi(popStack())
			pushStack(strconv.Itoa(v2 - v1))
		case "*":
			v1, _ := strconv.Atoi(popStack())
			v2, _ := strconv.Atoi(popStack())
			pushStack(strconv.Itoa(v1 * v2))
		case "/":
			v1, _ := strconv.Atoi(popStack())
			v2, _ := strconv.Atoi(popStack())
			pushStack(strconv.Itoa(v2 / v1))
		default:
			pushStack(v)
		}
	}

	result, _ := strconv.Atoi(stack[0])
	return result

}

func Test_evalRPN(t *testing.T) {
	a := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(a))
}
