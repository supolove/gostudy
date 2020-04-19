package 携程

import (
	"fmt"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	// 如果设置单线程，则第一个协程无法让出时间片
	// 第二个协程一直得不到时间片，阻塞等待。
	//runtime.GOMAXPROCS(2)

	flag := true

	go func() {
		fmt.Printf("coroutine one start \n")
		i := 0
		for flag {
			i++
			// 如果加了这行代码，协程可以让时间片
			// 这个因为 fmt.Printf 是内联函数，这是种特殊情况
			fmt.Printf("i: %d \n", i)
			// 挂起改协程
			//runtime.Gosched()
		}
		fmt.Printf("coroutine one exit \n")
	}()

	go func() {
		fmt.Printf("coroutine two start \n")
		flag = false
		fmt.Printf("coroutine two exit \n")
	}()

	time.Sleep(5 * time.Second)
	fmt.Printf("end \n")
}
