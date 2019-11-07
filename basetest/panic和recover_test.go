package basetest

import (
	"fmt"
	"testing"
)

func TestPanicAndRecover(t *testing.T) {
	//defer所在函数执行完所有的代码之后，会自动执行defer的这个函数，可以说是析构函数，定义方式：defer function_name()

	//给一个例子，获取素组元素，处理素组访问越界的问题。
	a := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 10; i++ {
		item, ok := get(i, a)
		fmt.Printf("a[%d]=%d[%v]\n", i, item, ok)
	}

	// recover 相当于try-catch的catch部分， 使得panic不在传递。而defer相当于try-catch-final的final部分。
	defer func() {
		fmt.Println("a")
		if err := recover(); err != nil {
			fmt.Println("panic info is ", err)
		}
	}()

	fmt.Println("b")
	panic("test function t")
	fmt.Println("c")

	// panic要点：panic相当于一个运行时异常，遇到panic的时候，会停止当前函数剩下来的语句，但在退出该函数之前，会执行defer的语句，依据函数调用层次，panic依次终止每个函数，直至main()。
	// recover 一般在defer中使用，用来捕获panic异常，当捕获到panic程序恢复正常运行
}

func get(i int, a [5]int) (ret int, ok bool) {
	ok = true
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "[set to default value -1 ]")
			ret = -1
			ok = false
		}
	}()

	ret = a[i]
	return
}
