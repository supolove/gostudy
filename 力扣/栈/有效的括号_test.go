package 栈

import (
	"fmt"
	"testing"
)

/**
说明：
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
  1.左括号必须用相同类型的右括号闭合。
  2.左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
  输入: "()"
  输出: true

示例 2:
  输入: "()[]{}"
  输出: true

示例 3:
  输入: "(]"
  输出: false

示例 4:
  输入: "([)]"
  输出: false

解题思路：
  1.把左括号压入栈
  2.遇到右边括号,退栈匹配元素
  3.把未匹配到的继续压入栈
  4.知道所有的匹配完成
*/

// 没看清楚题目，以为任意位置只要有闭合就是true
func isValid1(s string) bool {

	b := []byte(s)

	isLeft := func(ss byte) bool {
		if ss == '(' || ss == '{' || ss == '[' {
			return true
		}
		return false
	}

	if len(s) == 0 {
		return true
	}

	if !isLeft(b[0]) {
		return false
	}

	stack := make([]byte, 0)
	stackTemp := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		item := b[i]
		if isLeft(item) {
			stack = append(stack, b[i])
		} else {
			for len(stack) > 0 {
				var pop = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				if 0 < item-pop && item-pop < 3 {
					for len(stackTemp) > 0 {
						stack = append(stack, stackTemp[len(stackTemp)-1])
						stackTemp = stackTemp[0 : len(stackTemp)-1]
					}
					break
				} else {
					stackTemp = append(stackTemp, pop)
				}
			}

			if len(stack) == 0 {
				break
			}
		}
	}

	if len(stack) == 0 && len(stackTemp) == 0 {
		return true
	}

	return false
}

// 判断相邻的就可以了
func isValid(s string) bool {

	b := []byte(s)

	isLeft := func(ss byte) bool {
		if ss == '(' || ss == '{' || ss == '[' {
			return true
		}
		return false
	}

	if len(s) == 0 {
		return true
	}

	if !isLeft(b[0]) {
		return false
	}

	stack := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		item := b[i]
		if isLeft(item) {
			stack = append(stack, b[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			var pop = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if 0 < item-pop && item-pop < 3 {
				continue
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func Test_IsValid(t *testing.T) {
	fmt.Println('(', ')'-1)
	fmt.Println('[', ']')
	fmt.Println('{', '}')

	fmt.Println(isValid("[])"))
}
